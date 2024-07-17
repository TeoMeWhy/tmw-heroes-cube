FROM python:3.12.4

WORKDIR /app

COPY . .

RUN pip install -r requirements.txt

EXPOSE 8501

CMD ["streamlit", "run", "Heroes_of_Cube.py"]