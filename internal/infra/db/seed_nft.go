package db

import (
	"context"

	"github.com/orbit-alliance/orbit-backend/internal/domain/nft"
)

func SeedNFTs(ctx context.Context, nftRepo *NFTRepository) {

	if len(nftRepo.LoadAll(ctx)) > 0 {
		return
	}

	nftMetadata := nft.NewNFTMetadata("image_url", nil)

	nfts := []nft.NFT{
		*nft.NewNFT("Zelador", "Reconhece o cadete como alguém comprometido com a rotina do campus. Associado ao primeiro grande marco de presença contínua.", *nftMetadata),
	}

	for _, nft := range nfts {
		err := nftRepo.Save(ctx, &nft)
		if err != nil {
			panic(err)
		}
	}

}
