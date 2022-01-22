package sam.henhaochi.authservice.controllers.mappers;

import org.springframework.stereotype.Service;
import sam.henhaochi.authservice.controllers.requests.RegisterAccountRequest;
import sam.henhaochi.authservice.entities.Account;

@Service
public class RegisterAccountRequestMapper {
    public Account mapToAccountEntity(RegisterAccountRequest request){
        return Account.builder()
                .username(request.username)
                .password(request.password)
                .email(request.email)
                .build();
    }
}
