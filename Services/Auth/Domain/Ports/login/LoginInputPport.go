package auth_domain_ports

import auth_domain_dto "delivery/Services/Auth/Domain/DTOs"

type LoginInputPort interface {
	Login(dto auth_domain_dto.LoginDTO) (string, error)
}
