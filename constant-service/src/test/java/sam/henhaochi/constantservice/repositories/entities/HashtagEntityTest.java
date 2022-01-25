package sam.henhaochi.constantservice.repositories.entities;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class HashtagEntityTest {
    private static String ID = "";
    private static String NAME = "";
    private static int COUNT = 0;

    @Test
    public void shouldReturnIncrementedCount(){
        HashtagEntity result = HashtagEntity.builder()
                .id(ID)
                .name(NAME)
                .count(COUNT)
                .build();

        assertEquals(
                COUNT + 1,
                result.incrementCount()
        );
    }
}