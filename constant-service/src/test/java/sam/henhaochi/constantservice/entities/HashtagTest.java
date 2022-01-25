package sam.henhaochi.constantservice.entities;

import org.junit.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;

public class HashtagTest {

    private static String ID = "";
    private static String NAME = "";
    private static int COUNT = 0;

    @Test
    public void shouldCreateHashtagWithSuppliedValue(){
        Hashtag result = Hashtag.builder()
                .id(ID)
                .name(NAME)
                .count(COUNT)
                .build();

        assertEquals(
                ID,
                result.getId()
        );
        assertEquals(
                NAME,
                result.getName()
        );
        assertEquals(
                COUNT,
                result.getCount()
        );
    }
}
