package sam.henhaochi.constantservice.controllers.structure;

import lombok.Builder;
import lombok.Data;
import lombok.Getter;
import lombok.Setter;

@Data
@Builder
@Getter
@Setter
public class HashtagResponse {
    private String id;
    private String name;
    private int count;
}
