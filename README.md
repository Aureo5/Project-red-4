Project-red-4
Project red groupe 4
=======
Project red groupe 4
La bataille d'Osmos

---notre jeu---

Dans ce jeu vous aurez la possibilité de choisir votre nom, une classe avec des compétences différentes et l'orientation de vos choix. Que ce soit dans le menu, dans les différents districts ou  dans le système de combat tour par tour vous serez toujours maître de vos choix. 


---Aperçu---

<img width="64" height="63" alt="image" src="https://github.com/user-attachments/assets/e96bb119-a6e2-44a2-9531-af34942ae392" />

Ci-dessus notre menu qui vous permettras de lancer une partie.


<img width="683" height="251" alt="image" src="https://github.com/user-attachments/assets/b9a94549-e895-400f-b2a1-4fc70a8590bb" />

Ensuite, nous aurons l'Histoire qui permettra de faire vivre ce jeu puis le choix de votre nom et des différentes classes proposées.


<img width="492" height="151" alt="image" src="https://github.com/user-attachments/assets/30fbf1f4-fb44-4263-a4c1-3f8cdba8ff68" />

Et enfin, comme vu précédemment, le choix de l'orientation : combattre, allez au camp (marchand,forgeron...), aller au prochain district et accéder à votre inventaire.


---Arborescence---

.
├── Camp/
│   ├── firstcamp.go
│   ├── lastcamp.go
│   └── secondcamp.go
│
├── Character/
│   ├── Additem/
│   │   └── Additem.go
│   │
│   ├── class/
│   │   └── class.go
│   │
│   ├── inventory/
│   │   ├── fight-inventory/
│   │   │   └── fightinventory.go
│   │   ├── inventoryInteractive/
│   │   │   └── InventoryInteractive.go
│   │   └── inventorydisplay/
│   │       └── inventory.go
│   │
│   └── character.go
│
├── Npc/
│   ├── Blacksmith/
│   │   └── Blacksmith.go
│   │
│   ├── Enemy/
│   │   ├── beartitan.go
│   │   ├── coloss.go
│   │   ├── colossaletitan.go
│   │   ├── giant.go
│   │   ├── loot.go
│   │   ├── pattern.go
│   │   └── wolftitan.go
│   │
│   ├── Lastblacksmith/
│   │   └── lastblacksmith.go
│   │
│   ├── Merchant/
│   │   └── Merchant.go
│   │
│   └── Secblacksmith/
│       └── secblacksmith.go
│
├── district/
│   ├── endgame/
│   │   └── endgame.go
│   │
│   ├── firstdistrict/
│   │   └── firstdistrict.go
│   │
│   ├── lastdistrict/
│   │   └── lastdistrict.go
│   │
│   └── seconddistrict/
│       └── seconddistrict.go
│
├── equipment/
│   ├── equip/
│   │   └── equip.go
│   │
│   ├── removeitem/
│   │
│   ├── Propulsor.go
│   ├── bearsclaw.go
│   ├── bearsgrapplinghook.go
│   ├── bearspropulsor.go
│   ├── equipment.go
│   ├── grapplinghook.go
│   ├── lame.go
│   ├── wolfgrapplinghook.go
│   ├── wolfpropulsor.go
│   └── wolfsfang.go
│
├── fight/
│   ├── bossfight/
│   │   └── bossfight.go
│   │
│   ├── randomfight/
│   │   └── randomfight.go
│   │
│   ├── training/
│   │   └── trainingfight.go
│   │
│   └── fight.go
│
├── item/
│   ├── bear-skin.go
│   ├── bone.go
│   ├── combatsyringe.go
│   ├── grenadeincendiaire.go
│   ├── healsyringe.go
│   ├── item.go
│   ├── skin.go
│   └── wolf-skin.go
│
├── lore/
│   └── lore.go
│
├── menu/
│   └── menu.go
│
├── skill/
│   └── skill.go
│
├── README.md
├── go.mod
└── main.go


---Prérequis---

Pour le bon fonctionnement du jeux, il vous faut installer : Visual Studio Code version 1.27.1; Golang version 1.27.1



---Installation---

Installez les fichiers en zip, extraire le contenue dans un fichier de votre choix (de préférence un nouveau vide).



---Utilisation---

<img width="267" height="140" alt="image" src="https://github.com/user-attachments/assets/f50c48e6-f5bf-4185-a3ea-683c56dfb0e4" />


Cette fonction sert à lancer notre jeu.
Avant de lancer le jeu il faut que tu sois dans le document "Project-red-4", pour cela il faut faire dans le terminal : 
cd Project-red-4.
Ensuite pour lancer le jeu il faut écrire : "go run ." dans le terminal.



---Contribution---

Pour pouvoir contribuer au projet : 
A écrire dans le terminal :
1. se connecter à notre repository -> "git clone https://github.com/Aureo5/Project-red-4.git" dans ton terminal
2. Créer une branche -> git branch <nom de ta branche>
3. Accéder à ta branche -> git checkout <nom de ta branche crée>

Une fois tes modifications faites :  
1. Choisir tout les fichiers -> git add .
2. Enregistrer les changements -> git commit -m "<nom choisi>"
3. Envoyer vers ta branche sur github -> git push

Une fois tes modifications faites et prête à être mise sur la branche principal (main) :
1. Aller sur la branche "main" -> git checkout main
2. Récupérer la dernière version de "main" -> git pull origin main
3. Fusionner ta branche dedans -> git merge feature/<nom de ta branche>
4. Envoyer le résultat vers GitHub -> git push origin main


>>>>>>> b34673b02436bdb804a060e47ac13e5bf2c67a81
