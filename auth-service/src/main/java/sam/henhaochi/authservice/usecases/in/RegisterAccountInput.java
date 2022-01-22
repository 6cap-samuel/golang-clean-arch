package sam.henhaochi.authservice.usecases.in;

import sam.henhaochi.authservice.constants.AccountCreationStatus;
import sam.henhaochi.authservice.entities.Account;

import java.security.NoSuchAlgorithmException;

public interface RegisterAccountInput {
    AccountCreationStatus with(
            Account account
    ) throws NoSuchAlgorithmException;
}
