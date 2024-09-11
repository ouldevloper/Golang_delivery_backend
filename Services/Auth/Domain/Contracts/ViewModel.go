package auth_domain_contracts

import shared_models "delivery/Services/Shared/Infrastructure/Models"

type ViewModel interface {
	GetResponse() shared_models.ResponseModel
	String() string
}
