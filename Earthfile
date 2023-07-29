VERSION 0.7

FROM tochemey/docker-go:1.20.4-0.8.0

protogen:
    # copy the proto files to generate
    COPY --dir protos/ ./
    COPY buf.work.yaml buf.gen.yaml ./

    # generate the pbs
    RUN buf generate \
            --template buf.gen.yaml \
            --path protos/chief-of-state-protos/chief_of_state/v1

    # save artifact to
    SAVE ARTIFACT gen gen AS LOCAL gen