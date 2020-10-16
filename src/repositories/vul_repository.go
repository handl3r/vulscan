package repositories

type VulRepository struct {
	baseRepository
}

func NewVulRepository(baseRepository baseRepository) *VulRepository {
	return &VulRepository{baseRepository: baseRepository}
}
