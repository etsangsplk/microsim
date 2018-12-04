package main

import (
	"github.com/yurishkuro/microsim/model"
)

var hotrod = model.Config{
	Services: []*model.Service{
		&model.Service{
			Name: "ui",
		},
		&model.Service{
			Name: "frontend",
			Depends: &model.Dependencies{
				Seq: model.Sequence{
					{Service: &model.ServiceDep{Name: "customer"}},
					{Service: &model.ServiceDep{Name: "driver"}},
					{Service: &model.ServiceDep{Name: "route"}},
				},
			},
		},
		&model.Service{
			Name: "customer",
		},
		&model.Service{
			Name:  "driver",
			Count: 2,
		},
		&model.Service{
			Name:  "route",
			Count: 3,
		},
		&model.Service{
			Name: "mysql",
		},
		&model.Service{
			Name: "redis",
		},
	},
}