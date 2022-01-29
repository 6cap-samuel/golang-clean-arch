package sam.henhaochi.authservice.controllers.mappers;

import org.springframework.stereotype.Service;
import sam.henhaochi.authservice.controllers.responses.LoginResponse;
import sam.henhaochi.authservice.entities.Account;

@Service
public class LoginAccountResponseMapper {
    public LoginResponse mapToResponseEntity(Account account){
        return LoginResponse.builder()
                .role(account.getRole())
                .build();
    }
}
