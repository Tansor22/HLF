### Поднятие сети с чейнкодом
1. Запуск виртуальной машины и логин в ней  
`vagrant up`  
`vagrant ssh`  
2. Поднять сеть скриптом`dev-init.ssh -s`
3. Развернуть код на пире, пример в скрипте`network/bin/deploy-erc20.sh`
### Вызов чейнкода извне
1. Все обращения делаются от лица какого-нибудь юзера, 
   в `middleware/gateway/gateway.js` от лица админа.
   
2. **!!!ВАЖНО!!!** (КАЖДЫЙ РАЗ ПРИ ПЕРЕЗАПУСКЕ СЕТИ НУЖНО ГЕНЕРИРОВАТЬ НОВЫЕ СЕРТИФИКАТЫ) Нужно чтобы в `middleware/gateway/user-wallet` были сертификаты юзера,
   от которого идет обращение к чейнкоду. 
   Для этого нужно генерить сертификаты утилитой `middleware/gateway/wallet.js`.
