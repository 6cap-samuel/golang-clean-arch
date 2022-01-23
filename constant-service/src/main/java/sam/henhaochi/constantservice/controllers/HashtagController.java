package sam.henhaochi.constantservice.controllers;

import lombok.AllArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import sam.henhaochi.constantservice.controllers.mappers.HashtagStructureMapper;
import sam.henhaochi.constantservice.controllers.responses.RetrieveHashtagsResponse;
import sam.henhaochi.constantservice.controllers.structure.HashtagResponse;
import sam.henhaochi.constantservice.usecases.in.RetrieveHashtagsInput;

import java.util.List;
import java.util.stream.Collectors;

@AllArgsConstructor
@RestController
@RequestMapping("/hashtags")
public class HashtagController {

    final HashtagStructureMapper mapper;
    final RetrieveHashtagsInput retrieve;

    private static final Logger logger = LoggerFactory.getLogger(HashtagController.class);

    @GetMapping
    public ResponseEntity<RetrieveHashtagsResponse> get(){
        logger.info("GET: /hashtags called");
        List<HashtagResponse> responses = retrieve.get()
                .stream()
                .map(mapper::map)
                .collect(Collectors.toList());

        return ResponseEntity.ok(
                RetrieveHashtagsResponse.builder()
                        .hashtags(responses)
                        .build()
        );
    }
}
