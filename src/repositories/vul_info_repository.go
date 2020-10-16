package repositories

type VulInfoRepository struct {
	baseRepository
}

func NewVulInfoRepository(baseRepository baseRepository) *VulInfoRepository {
	return &VulInfoRepository{baseRepository: baseRepository}
}
