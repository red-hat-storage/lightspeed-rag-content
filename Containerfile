# Tool image: embedding model and generate_embeddings_tool.py.
# https://docs.redhat.com/en/documentation/red_hat_openshift_lightspeed/1.0/html/configure/ols-configuring-openshift-lightspeed#providing-custom-knowledge-to-the-llm_ols-configuring-openshift-lightspeed
ARG BYOK_TOOL_IMAGE=quay.io/redhat-user-workloads/crt-nshift-lightspeed-tenant/lightspeed-rag-tool:latest
# Final image. The operator mounts /rag/vector_db from this stage.
ARG UBI_BASE_IMAGE=registry.access.redhat.com/ubi9/ubi-micro@sha256:f45ee3d1f8ea8cd490298769daac2ac61da902e83715186145ac2e65322ddfc8

# STAGE 1: Embed the parsed Markdown. Each file needs title and url front matter.
FROM ${BYOK_TOOL_IMAGE} as tool
ARG DOC_VERSION
USER 0
WORKDIR /workdir
COPY parsed/${DOC_VERSION}/ /markdown/
COPY odf/dr/recipes/markdown/ /markdown/recipes/
RUN python3.12 generate_embeddings_tool.py \
    -i /markdown \
    -emd embeddings_model \
    -emn sentence-transformers/all-mpnet-base-v2 \
    -o vector_db

# STAGE 2: Ship only the vector database.
FROM ${UBI_BASE_IMAGE}
COPY --from=tool /workdir/vector_db /rag/vector_db
