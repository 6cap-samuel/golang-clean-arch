package sam.henhaochi.constantservice.controllers;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import sam.henhaochi.constantservice.controllers.mappers.HashtagStructureMapper;
import sam.henhaochi.constantservice.controllers.structure.HashtagResponse;
import sam.henhaochi.constantservice.entities.Hashtag;
import sam.henhaochi.constantservice.usecases.in.RetrieveHashtagsInput;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;

import java.util.List;

import static org.mockito.Mockito.when;

import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

@WebMvcTest(HashtagController.class)
class HashtagControllerTest {
    private static String ID = "";
    private static String NAME = "";
    private static int COUNT = 0;

    @Autowired
    private MockMvc mockMvc;

    @MockBean
    private HashtagStructureMapper mapper;

    @MockBean
    private RetrieveHashtagsInput retrieve;

    @Test
    public void shouldReturnHashtagResultWhenGetIsCalled()
            throws Exception {
        Hashtag hashtag = Hashtag.builder()
                .id(ID)
                .name(NAME)
                .count(COUNT)
                .build();
        List<Hashtag> hashtags = List.of(hashtag);

        when(retrieve.get()).thenReturn(hashtags);
        when(mapper.map(hashtag)).thenReturn(
                HashtagResponse.builder()
                        .id(ID)
                        .name(NAME)
                        .count(COUNT)
                        .build()
        );

        mockMvc.perform(get("/hashtags"))
                .andExpect(status()
                        .isOk()
                ).andExpect(content()
                        .contentType(
                                MediaType.APPLICATION_JSON
                        )
                ).andExpect(
                        jsonPath("$.hashtags")
                                .exists()
                ).andExpect(
                        jsonPath("$.hashtags[0].id")
                                .value(ID)
                ).andExpect(
                        jsonPath("$.hashtags[0].name")
                                .value(NAME)
                ).andExpect(
                        jsonPath("$.hashtags[0].count")
                                .value(COUNT)
                );
    }
}