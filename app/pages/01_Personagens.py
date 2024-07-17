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

def format_slots(slots):

    slots = slots.T
    
    df_slots = pd.DataFrame(index=["foot", "head", "chest", "weapon"])

    slots = pd.concat([df_slots,slots], axis=1)

    slots = (slots[['Id', 'Name', 'Damage', 'Price', 'Class']]
                  .rename(columns={
                        'Name':'Nome',
                        'Damage':'Dano',
                        'Price':'Preço',
                        'Class': 'Classe',
                        })
    )

    slots.index = ["👢", "🎩", "👕" ,"⚔️"]

    return slots


st.markdown("# Personagens!")

st.text_input("Seu nick na Twitch", key="name", help="Insira seu nick da Twitch para buscarmos seu personagem.")
name = st.session_state.name.lower()

if name != "":
    
    resp = requests.get(f"http://heroes:8085/persons/?name={name}")

    if resp.status_code == 200:
        data = resp.json()
        id_person = data["id"]
        
        url_points = f"http://upsell:8080/customers?id={id_person}"
        resp_points = requests.get(url=url_points)
        data['points'] = resp_points.json()["Points"]

        text = format_player(data)
        st.markdown(text)


        st.markdown("""
        ---
        ## Equipamentos""")

        slots_url = f"http://heroes:8085/slots/{id_person}"
        resp_slots = requests.get(slots_url)
        slots = resp_slots.json()
        slots = pd.DataFrame(slots)
        if slots.shape[0] != 0:
            slots = format_slots(slots)
        
        st.dataframe(slots)

        st.markdown("""
Para equipar itens a partir do seu inventário, use: `!equipe id`
Assim, para equipar suas botas que são id=1:
        
        !equip 1
                            
Para devolver um item equipado ao inventário, use `!unequip id`.
Assim, para desequipar suas botas que são id=1:
                    
        !unequip 1
                        
        """)


        st.markdown("""
                    ---
                    ## Inventário""")

        resp_inventory = requests.get(f"http://heroes:8085/inventories/{id_person}")
        inventory = pd.DataFrame(resp_inventory.json())
        inventory = format_inventory(inventory=inventory)
        
        col1, col2 = st.columns([2,1])

        with col1:
            st.dataframe(inventory, hide_index=True)
            
            txt = """
Para vender seus itens do inventário, use: `!sell id`

Ou seja, para vender seu item de `id=1`, use:
            
            !sell 1
            
            """
            
            st.markdown(txt)

        with col2:
            url = "https://www.twitch.tv/embed/teomewhy/chat?parent=emu-blessed-incredibly.ngrok-free.app"
            components.iframe(src=url,
                              height=400,
                              width=300)
    else:
        st.error("Personagem não encontrado")

