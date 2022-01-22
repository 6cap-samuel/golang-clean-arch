package sam.henhaochi.authservice.controllers.mappers;

import org.springframework.stereotype.Service;
import sam.henhaochi.authservice.controllers.requests.LoginRequest;
import sam.henhaochi.authservice.entities.Account;

@Service
public class LoginAccountRequestMapper {
    public Account mapToAccountEntity(LoginRequest request){
        return Account.builder()
                .username(request.username)
                .password(request.password)
                .email(request.email)
                .build();
    }
}
