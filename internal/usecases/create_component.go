package usecases

import "algvisual/internal/database"

type CreateComponenteUseCaseRequestBody struct {
	elements_id  []int
	color        string
	component_id string
}

func CreateComponentUseCase(db database.DBTX, req CreateComponenteUseCaseRequestBody) {

}
