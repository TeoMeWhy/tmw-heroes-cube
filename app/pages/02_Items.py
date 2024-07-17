import streamlit as st
import pandas as pd
import requests

def format_items(df):
    df = df.T[['Id', 'Name', 'Damage', 'Price', "Class", "Type"]]
    df["Id"] = df["Id"].astype(int)
    df = (df.rename(columns={
                    'Name': "Nome",
                    'Damage': "Dano",
                    'Price': "Preço",
                    "Class": "Classe",
                    "Type": "Posição",})
            .sort_values(by='Id'))
    
    return df

def filter_df(df,  classe=None, posicao=None, nome=None):
    df = df.copy()
    if classe:
        df = df[df['Classe']==classe]

    if posicao:
        df = df[df['Posição']==posicao]

    if nome:
        df = df[df['Nome']==nome]

    return df


url = "http://heroes:8085/items"

resp = requests.get(url)
items = pd.DataFrame(resp.json())
items = format_items(items)

classes = items["Classe"].unique().tolist()
posicoes = items["Posição"].unique().tolist()
nomes = items["Nome"].unique().tolist()

st.markdown("# Items")

col1, col2 = st.columns([1,2])

with col1:
    classe = st.selectbox(label="Classe", options=classes, index=None, placeholder="Selecione uma classe")
    posicao = st.selectbox(label="Posição", options=posicoes, index=None, placeholder="Selecione uma posição do corpo")
    nome = st.selectbox(label="Nome", options=nomes, index=None, placeholder="Entre com o nome do item")

   
with col2:
    item_filer = filter_df(items, classe=classe, posicao=posicao, nome=nome)
    st.dataframe(item_filer, hide_index=True)
