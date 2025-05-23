// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"

	"github.com/absmach/magistrala/provision"
	apiutil "github.com/absmach/supermq/api/http/util"
	"github.com/absmach/supermq/pkg/errors"
	"github.com/go-kit/kit/endpoint"
)

func doProvision(svc provision.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(provisionReq)
		if err := req.validate(); err != nil {
			return nil, errors.Wrap(apiutil.ErrValidation, err)
		}

		res, err := svc.Provision(ctx, req.domainID, req.token, req.Name, req.ExternalID, req.ExternalKey)
		if err != nil {
			return nil, err
		}

		provisionResponse := provisionRes{
			Clients:     res.Clients,
			Channels:    res.Channels,
			ClientCert:  res.ClientCert,
			ClientKey:   res.ClientKey,
			CACert:      res.CACert,
			Whitelisted: res.Whitelisted,
		}

		return provisionResponse, nil
	}
}

func getMapping(svc provision.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(mappingReq)
		if err := req.validate(); err != nil {
			return nil, errors.Wrap(apiutil.ErrValidation, err)
		}

		res, err := svc.Mapping(ctx, req.token)
		if err != nil {
			return nil, err
		}

		return mappingRes{Data: res}, nil
	}
}
