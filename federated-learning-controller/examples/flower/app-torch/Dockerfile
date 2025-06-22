# Use a smaller base image
FROM python:3.11-slim

# Set environment variables to reduce image size and improve pip behavior
ENV PYTHONUNBUFFERED=1 \
  PIP_NO_CACHE_DIR=1 \
  PIP_DISABLE_PIP_VERSION_CHECK=1 \
  PIP_DEFAULT_TIMEOUT=100 \
  MPLCONFIGDIR=/tmp \
  HF_HOME=/app/.cache/huggingface \
  TRANSFORMERS_CACHE=/app/.cache/transformers

# Set the working directory
WORKDIR /app

# Copy only necessary files (excluding pyproject.toml to defer dependency installation)
COPY . /app

# Ensure the cache directories have the correct permissions
RUN chown -R root:root /app && chmod -R 777 /app


# Set the PYTHONPATH environment variable
ENV PYTHONPATH=/app

# Ensure the entrypoint script is executable
RUN chmod +x /app/entrypoint.sh

ENV HOME=/app
ENV PIP_NO_CACHE_DIR=1
ENV VENV_PATH=/app/.venv

RUN mkdir -p $VENV_PATH && chmod -R 777 $VENV_PATH

RUN pip install --no-cache-dir --prefix=$VENV_PATH -e . -i https://mirrors.ustc.edu.cn/pypi/simple

ENV http_proxy="http://172.16.1.11:20172"
ENV https_proxy="http://172.16.1.11:20172"
ENV no_proxy="localhost,127.0.0.1,::1,10.0.0.0/8,172.16.0.0/12,192.168.0.0/16"

# Install dependencies at runtime
# ENTRYPOINT ["sh", "-c", "pip install -e . && /app/entrypoint.sh"]
ENTRYPOINT ["/app/entrypoint.sh"]
