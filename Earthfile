VERSION 0.8

FROM tochemey/docker-go:1.22.5-3.2.0

protogen:
    # copy the proto files to generate
    COPY --dir protos/ ./
    COPY buf.yaml buf.gen.yaml ./

    # generate the pbs
    RUN buf generate \
            --template buf.gen.yaml \
            --path protos/chief-of-state-protos/chief_of_state/v1

    # save artifact to
    SAVE ARTIFACT gen gen AS LOCAL gen