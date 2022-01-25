package sam.henhaochi.constantservice.repositories.entities;

import lombok.Builder;
import lombok.Data;
import lombok.Getter;
import lombok.Setter;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Data
@Builder
@Getter
@Setter
@Document("hashtags")
public class HashtagEntity {
    @Id
    private String id;
    private String name;
    private int count;

    public int incrementCount(){
        return ++count;
    }
}

