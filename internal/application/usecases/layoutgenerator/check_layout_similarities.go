package layoutgenerator

import (
	"algvisual/internal/application/usecases/designassets"
	"algvisual/internal/application/usecases/geometry"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"context"
	"math"

	"github.com/jackc/pgx/v5/pgtype"
)

type CheckLayoutSimilaritiesInput struct {
	RequestID int32
	Layout    entities.Layout
}

type CheckLayoutSimilaritiesOutput struct {
	HaveSimilar bool
}

func CheckLayoutSimilaritiesUseCase(
	ctx context.Context,
	req CheckLayoutSimilaritiesInput,
	db *database.Queries,
	das *designassets.DesignAssetService,
) (*CheckLayoutSimilaritiesOutput, error) {
	layouts, err := db.GetLayoutByRequestID(ctx, pgtype.Int4{Int32: req.RequestID, Valid: true})
	if err != nil {
		return nil, err
	}
	for _, l := range layouts {
		data, err := GetLayoutByIDUseCase(ctx, db, GetLayoutByIDRequest{
			LayoutID: int32(l.ID),
		}, das)
		if len(data.Layout.Components) != len(req.Layout.Components) {
			continue
		}
		if err != nil {
			return nil, err
		}
		isEqual := true
		for _, comp := range data.Layout.Components {
			for _, el := range comp.Elements {
				for _, reqElement := range req.Layout.Elements {
					if reqElement.AssetID != el.AssetID {
						continue
					}
					wdf := float64(reqElement.Width()) - float64(el.Width())
					if math.Abs(wdf) > 20 {
						isEqual = false
						continue
					}
					hdf := float64(reqElement.Height()) - float64(el.Height())
					if math.Abs(hdf) > 20 {
						isEqual = false
						continue
					}
					isSimilar := geometry.IsContainerSimilar(
						reqElement.OuterContainer,
						el.OuterContainer,
						20,
					)
					if !isSimilar {
						isEqual = false
					}
				}
			}
		}
		if isEqual {
			return &CheckLayoutSimilaritiesOutput{HaveSimilar: true}, nil
		}
	}
	return &CheckLayoutSimilaritiesOutput{HaveSimilar: false}, nil
}

func IsSimilar(base entities.Layout, compareTo entities.Layout) bool {
	isEqual := true
	for _, comp := range compareTo.Components {
		for _, el := range comp.Elements {
			for _, reqElement := range base.Elements {
				if reqElement.AssetID != el.AssetID {
					continue
				}
				wdf := float64(reqElement.Width()) - float64(el.Width())
				if math.Abs(wdf) > 20 {
					isEqual = false
					continue
				}
				hdf := float64(reqElement.Height()) - float64(el.Height())
				if math.Abs(hdf) > 20 {
					isEqual = false
					continue
				}
				isSimilar := geometry.IsContainerSimilar(
					reqElement.OuterContainer,
					el.OuterContainer,
					20,
				)
				if !isSimilar {
					isEqual = false
				}
			}
		}
	}
	return isEqual
}
