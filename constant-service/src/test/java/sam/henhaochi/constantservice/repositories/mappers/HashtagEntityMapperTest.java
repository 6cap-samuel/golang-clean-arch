package sam.henhaochi.constantservice.repositories.mappers;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import sam.henhaochi.constantservice.entities.Hashtag;
import sam.henhaochi.constantservice.repositories.entities.HashtagEntity;

import static org.junit.jupiter.api.Assertions.*;

class HashtagEntityMapperTest {

    private static final String NAME = "#hashtag";
    private static final int COUNT = 1;
    private static final String ID = "id";

    private HashtagEntityMapper hashtagEntityMapper;

    @BeforeEach
    public void init() {
        hashtagEntityMapper = new HashtagEntityMapper();
    }

    @Test
    public void shouldMapHashtagNameToEntity(){
        HashtagEntity hashtag =
                hashtagEntityMapper.mapStringToEntity(NAME);

        assertNotNull(hashtag.getId());
        assertEquals(hashtag.getName(), NAME);
        assertEquals(hashtag.getCount(), 1);
    }

    @Test
    public void shouldThrowExceptionWhenHashtagNameIsNull() {
        assertThrows(NullPointerException.class, () -> {
            hashtagEntityMapper.mapStringToEntity(null);
        });
    }

    @Test
    public void shouldMapHashtagEntityToHashtag(){
        Hashtag hashtag =
                hashtagEntityMapper.map(
                        HashtagEntity.builder()
                                .name(NAME)
                                .id(ID)
                                .count(COUNT)
                                .build()
                );

        assertEquals(hashtag.getName(), NAME);
        assertEquals(hashtag.getCount(), COUNT);
        assertEquals(hashtag.getId(), ID);
    }
}