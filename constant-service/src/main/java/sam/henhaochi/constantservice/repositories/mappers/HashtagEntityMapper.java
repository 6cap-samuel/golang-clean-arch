package sam.henhaochi.constantservice.repositories.mappers;

import org.springframework.stereotype.Service;
import sam.henhaochi.constantservice.entities.Hashtag;
import sam.henhaochi.constantservice.repositories.entities.HashtagEntity;

import java.util.UUID;

@Service
public class HashtagEntityMapper {
    public HashtagEntity mapStringToEntity(
            String name
    ) {
        return HashtagEntity.builder()
                .id(UUID.randomUUID().toString())
                .count(1)
                .name(name)
                .build();
    }

    public Hashtag map(HashtagEntity hashtagEntity){
        return Hashtag.builder()
                .id(hashtagEntity.getId())
                .count(hashtagEntity.getCount())
                .name(hashtagEntity.getName())
                .build();
    }
}
