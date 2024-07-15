import streamlit as st
import streamlit.components.v1 as components

import requests
import pandas as pd

def format_player(player):

    text = f"""
    ### Nome: {player['name']} | Cubos: {player['points']}
    
    #### Raça: {player['race']} | Classe: {player['class']}
    
    #### Dano: {player['damage']} | Level: {player['level']}"""

    return text


def format_inventory(inventory):

    if inventory.shape[0] == 0:
        return pd.DataFrame()

    columns = ["Id", "Name", "Damage", "Price", "Class"]
    inventory = (
    inventory.groupby(columns)["Quantity"]
            .sum()
            .reset_index()
            .rename(columns={
                    "Name": "Nome",
                    "Damage": "Dano",
                    "Price": "Preço",
                    "Class": "Classe",
                    "Quantity": "Quantidade",
                }))
    inventory["Id"] = inventory["Id"].astype(int)
    inventory = inventory.sort_values(["Id"])
    return inventory


st.markdown("# Personagens!")

st.text_input("Seu nick na Twitch", key="name", help="Insira seu nick da Twitch para buscarmos seu personagem.")
name = st.session_state.name.lower()

if name != "":
    
    resp = requests.get(f"http://localhost:8085/persons/?name={name}")

    if resp.status_code == 200:
        data = resp.json()
        id_person = data["id"]
        
        url_points = f"http://localhost:8080/customers?id={id_person}"
        resp_points = requests.get(url=url_points)
        data['points'] = resp_points.json()["Points"]

        text = format_player(data)
        st.markdown(text)

        resp_inventory = requests.get(f"http://localhost:8085/inventories/{id_person}")
        inventory = pd.DataFrame(resp_inventory.json())
        inventory = format_inventory(inventory=inventory)

        st.markdown("""
                    ---
                    ## Inventário""")

        col1, col2 = st.columns(2)

        with col1:
            st.dataframe(inventory, hide_index=True)
            
            txt = """
            Para vender seus itens do inventário, digite:
            
            `!sell id`.

            Ou seja, para vender seu item de `id=1`, utilize:

            `!sell 1`            
            """
            
            st.markdown(txt)

        with col2:
            url = "https://www.twitch.tv/embed/teomewhy/chat?parent=emu-blessed-incredibly.ngrok-free.app"
            components.iframe(src=url,
                              height=400,
                              width=300)
    else:
        st.error("Personagem não encontrado")

