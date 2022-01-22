package sam.henhaochi.authservice.usecases.out;

import sam.henhaochi.authservice.constants.AccountCreationStatus;
import sam.henhaochi.authservice.entities.Account;

public interface AccountDataSource {
    Account loginWith(
            Account account
    );
    AccountCreationStatus registerWith(
            Account account
    );
}
