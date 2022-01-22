package sam.henhaochi.constantservice.entities;

import lombok.Builder;
import lombok.Data;
import lombok.Getter;
import lombok.Setter;

@Data
@Builder
@Getter
@Setter
public class Hashtag {
    private String id;
    private String name;
    private int count;
}