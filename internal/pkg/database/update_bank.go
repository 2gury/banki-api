package database

import (
	"context"

	"banks-api/internal/pkg/model"
	"github.com/pkg/errors"
)

func (r *BanksRepository) UpdateBank(ctx context.Context, bank *model.Bank) error {
	var (
		updateBankColums = []string{
			bankExternalIDColumn,
			bankExternalLegacyIDColumn,
			bankNameColumn,
			bankLogoColumn,
			bankURLColumn,
			bankDataColumn,
			bankCreatedAtColumn,
			bankUpdatedAtColumn,
		}

		updateBankValues = []interface{}{
			bank.ExternalID,
			bank.ExternalLegacyID,
			bank.Name,
			bank.Logo,
			bank.URL,
			bank.BankData,
			bank.CreatedAt,
			bank.UpdatedAt,
		}
	)

	if bank.ID > 0 {
		updateBankColums = append(updateBankColums, bankIDColumn)
		updateBankValues = append(updateBankValues, bank.ID)
	}

	query := psql.
		Insert(banksTableName).
		Columns(updateBankColums...).
		Values(updateBankValues...).
		Suffix(`ON CONFLICT(id) DO UPDATE SET 
					external_id=excluded.external_id,
					external_legacy_id=excluded.external_legacy_id,
					name=excluded.name,
					logo=excluded.logo,
					url=excluded.url,
					bank_data=excluded.bank_data,
					updated_at=now()`)

	rawSQL, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "GetBanks.ToSql")
	}

	_, err = r.pool.Exec(ctx, rawSQL, args...)
	if err != nil {
		return errors.Wrap(err, "UpdateBank.Exec")
	}

	return nil
}