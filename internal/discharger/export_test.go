// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package discharger

import (
	"github.com/juju/simplekv"

	"github.com/cscstestbed/candid/idp"
	"github.com/cscstestbed/candid/internal/discharger/internal"
	"github.com/cscstestbed/candid/internal/identity"
	"github.com/cscstestbed/candid/store"
)

var NewIDPHandler = newIDPHandler

type LoginInfo loginInfo

func NewVisitCompleter(params identity.HandlerParams, kvstore simplekv.Store, store store.Store) idp.VisitCompleter {
	return &visitCompleter{
		params:        params,
		identityStore: internal.NewIdentityStore(kvstore, store),
		place:         &place{params.MeetingPlace},
	}
}
