import streamlit as st
import streamlit.components.v1 as components

st.markdown("# Boas vindas ao Heroes of Cube!")
st.write( f'<span style="font-size: 78px; line-height: 1">🧙‍♂️</span>', unsafe_allow_html=True, )
st.markdown("Se divirta com a gente nesse minigame de RPG")
st.markdown("## Participe")

text = """
1. Acesse o chat da Twitch: [twitch.tv/teomewhy](https://twitch.tv/teomewhy)
2. Participe do nosso sistema de pontos digitando: `!join`
3. Crei seu pesonagem: `!create`
    - Escolha uma raça e uma classe
    - Raças: dwarf | elf | hobbit | human | poney
    - Classes: mage | thief | warrior | bard | cleric
    - Exemplo: `!create elf mage`
4. Obtenha seus loots diários: `!loot`
5. Confira seu inventário: `!inventory`
"""

st.markdown(text)

components.iframe(src="https://www.twitch.tv/embed/teomewhy/chat?parent=flying-goldfish-novel.ngrok-free.app",
                  height=500,
                  width=300)
