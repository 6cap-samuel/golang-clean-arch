package sam.henhaochi.authservice.controllers.requests;

import lombok.Data;

@Data
public class LoginRequest {
    public String username;
    public String password;
    public String email;
}
