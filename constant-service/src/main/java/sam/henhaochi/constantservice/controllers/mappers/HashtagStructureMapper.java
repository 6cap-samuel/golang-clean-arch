package sam.henhaochi.constantservice.controllers.mappers;

import org.springframework.stereotype.Service;
import sam.henhaochi.constantservice.controllers.structure.HashtagResponse;
import sam.henhaochi.constantservice.entities.Hashtag;

@Service
public class HashtagStructureMapper {
    public HashtagResponse map(Hashtag entity){
        return HashtagResponse
                .builder()
                .id(entity.getId())
                .name(entity.getName())
                .count(entity.getCount())
                .build();
    }
}
