//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package postgres

import (
	"encoding/json"
	"github.com/go-pg/pg/orm"
	"github.com/insolar/observer/configuration"
	"github.com/insolar/observer/internal/app/observer"
	"github.com/insolar/observer/internal/app/observer/collecting"
	"github.com/insolar/observer/observability"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type GroupSchema struct {
	tableName struct{} `sql:"groups"`

	Ref            []byte `sql:",pk"`
	Title          string
	Goal           string
	Image          string
	GroupOwner     []byte
	TreasureHolder []byte
	Status         string
	State          []byte
}

func NewGroupStorage(obs *observability.Observability, db orm.DB) *GroupStorage {
	errorCounter := obs.Counter(prometheus.CounterOpts{
		Name: "observer_group_storage_error_counter",
		Help: "",
	})
	return &GroupStorage{
		log:          obs.Log(),
		errorCounter: errorCounter,
		db:           db,
	}
}

type GroupStorage struct {
	cfg          *configuration.Configuration
	log          *logrus.Logger
	errorCounter prometheus.Counter
	db           orm.DB
}

func (s *GroupStorage) Insert(model *observer.Group) error {
	if model == nil {
		s.log.Warnf("trying to insert nil group model")
		return nil
	}
	row := groupSchema(model)
	res, err := s.db.Model(row).
		OnConflict("DO NOTHING").
		Insert(row)

	if err != nil {
		return errors.Wrapf(err, "failed to insert group %v, %v", row, err.Error())
	}

	if res.RowsAffected() == 0 {
		s.errorCounter.Inc()
		s.log.WithField("group_row", row).
			Errorf("failed to insert group")
	}
	return nil
}

func (s *GroupStorage) Update(model *observer.GroupUpdate) error {
	if model == nil {
		s.log.Warnf("trying to apply nil update group model")
		return nil
	}

	var productType string

	switch model.ProductType {
	case observer.MerryGoRound:
		productType = "merry-go-round"
	case observer.Saving:
		productType = "saving"
	}

	res, err := s.db.Model(&GroupSchema{}).
		Where("state=?", model.PrevState.Bytes()).
		Set("image=?,goal=?", model.Image, model.Goal).
		Set("type=?", productType).
		Set("title=?", model.Title).
		Set("treasure_holder=?", model.Treasurer.Bytes()).
		Set("state=?", model.GroupState.Bytes()).
		Update()

	if err != nil {
		return errors.Wrapf(err, "failed to update group =%v", model)
	}

	if !model.Treasurer.IsEmpty() {
		res, err = s.db.Model(&UserGroupSchema{}).
			Where("user_ref=?", model.Treasurer.Bytes()).
			Set("role=?", "treasurer").
			Update()

		if err != nil {
			return errors.Wrapf(err, "failed to update user_group =%v", model)
		}

		if res.RowsAffected() == 0 {
			s.errorCounter.Inc()
			s.log.WithField("upd", model).Errorf("failed to update user_group")
		}
	}
	if model.Membership != nil {
		for _, membershipStr := range model.Membership {
			byt := []byte(membershipStr)
			var membership collecting.Membership
			if err := json.Unmarshal(byt, &membership); err != nil {
				return nil
			}
			var dbState string
			var dbRole string
			switch membership.MemberStatus {
			case collecting.StatusInvite:
				dbState = "invited"
			case collecting.StatusActive:
				dbState = "active"
			case collecting.StatusInactive:
				dbState = "inactive"
			}

			switch membership.MemberRole {
			case collecting.RoleChairMan:
				dbRole = "admin"
			case collecting.RoleTreasure:
				dbRole = "treasurer"
			case collecting.RoleMember:
				dbRole = "member"
			}

			count, err := s.db.Model(&UserGroupSchema{}).Where("group_ref=?", model.GroupReference.Bytes()).
				Where("user_ref=?", membership.MemberRef.Bytes()).Count()
			if count == 0 {

				row := &UserGroupSchema{
					UserRef:         membership.MemberRef.Bytes(),
					GroupRef:        model.GroupReference.Bytes(),
					Role:            "member",
					Status:          "invited",
					StatusTimestamp: model.Timestamp,
				}

				res, err := s.db.Model(row).
					OnConflict("DO NOTHING").
					Insert(row)

				if err != nil {
					return errors.Wrapf(err, "failed to insert user-group %v, %v", row, err.Error())
				}

				if res.RowsAffected() == 0 {
					s.errorCounter.Inc()
					s.log.WithField("user-group_row", row).
						Errorf("failed to insert user-group")
				}
			}

			res, err = s.db.Model(&UserGroupSchema{}).
				Where("group_ref=?", model.GroupReference.Bytes()).
				Where("user_ref=?", membership.MemberRef.Bytes()).
				Set("status=?", dbState).
				Set("role=?", dbRole).
				Update()

			if err != nil {
				return errors.Wrapf(err, "failed to update user_group =%v", model)
			}

			if res.RowsAffected() == 0 {
				s.errorCounter.Inc()
				s.log.WithField("upd", model).Errorf("failed to update user_group")
			}
		}

	}
	return nil
}

func groupSchema(model *observer.Group) *GroupSchema {
	return &GroupSchema{
		Ref:        model.Ref.Bytes(),
		Title:      model.Title,
		Goal:       model.Goal,
		GroupOwner: model.ChairMan.Bytes(),
		Image:      model.Image,
		Status:     model.Status,
		State:      model.State.Bytes(),
	}
}