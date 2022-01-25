package sam.henhaochi.constantservice.controllers.mappers;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import sam.henhaochi.constantservice.controllers.structure.HashtagResponse;
import sam.henhaochi.constantservice.entities.Hashtag;

import static org.junit.jupiter.api.Assertions.assertEquals;

class HashtagStructureMapperTest {

    private static final String NAME = "#hashtag";
    private static final int COUNT = 1;
    private static final String ID = "id";

    private HashtagStructureMapper hashtagStructureMapper;

    @BeforeEach
    public void init() {
        hashtagStructureMapper = new HashtagStructureMapper();
    }

    @Test
    public void shouldReturnNewHashtagResponseWhenPassingInValidEntity() {
        HashtagResponse response =
                hashtagStructureMapper.map(
                        Hashtag.builder()
                                .name(NAME)
                                .id(ID)
                                .count(COUNT)
                                .build()
        );

        assertEquals(response.getName(), NAME);
        assertEquals(response.getCount(), COUNT);
        assertEquals(response.getId(), ID);
    }
}