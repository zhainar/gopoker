<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col">
        <h4 class="alert-heading">
          Poker planning
        </h4>
        <br>
        <div v-if="ws === undefined" class="alert alert-success">
          <input v-model="my.name" type="text" class="form-control" placeholder="Enter your name">
          <br>
          <button class="btn btn-success" @click="connect">Connect</button>
        </div>
        <div v-else class="alert alert-success">
          <h4>You are successfully connected</h4>
          <br>
          <button class="btn btn-outline-success" @click="disconnect">Disconnect</button>
        </div>
        <br>
        <h6>Текущая задача</h6>
        <h3>{{ task.text ? task.text : 'Не указана' }}</h3>
        <br>
        <div class="poker-cards">
          <button v-for="card in cards" :key="card"
                  class="poker-card btn"
                  :class="{'btn-outline-primary': task.mark !== card, 'btn-primary': task.mark === card}"
                  @click="selectCard(card)">
            {{ card }}
          </button>
        </div>
        <hr>
        <br>
        <div class="users">
          <h4>Участники</h4>
          <ul v-if="users.length > 0" class="list-group">
            <li v-for="user in users" :key="user.ID" class="list-group-item">
              {{ user.name }}
              <span v-if="user.mark" class="badge badge-primary">{{ user.mark }}</span>
              <span v-else class="badge"
                    :class="{'bg-success': user.status === 'connected', 'bg-secondary': user.status !== 'connected'}">
                {{ user.status }}
              </span>
            </li>
          </ul>
          <div v-else class="alert alert-light">
            пусто
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'App',
  data () {
    return {
      ws: undefined,
      my: {
        name: ''
      },
      users: [],
      task: {
        text: '',
        mark: undefined
      },
      cards: [
        0.5, 1, 3, 5, 8, 13, 20, 40, 100, '?'
      ]
    }
  },
  methods: {
    connect () {
      if (this.ws) {
        return
      }

      const ws = new WebSocket(this.$settings.wsAddr)

      const el = this

      ws.onopen = function (ev) {
        el.ws = ws
        el.register()
      }

      ws.onmessage = this.handle
      ws.onerror = console.log
    },
    register (ev) {
      this.send('connect', {
        id: this.getID(),
        name: this.my.name ? this.my.name : ''
      })
    },
    disconnect () {
      this.send('disconnect')
    },
    async send (action, data) {
      if (!this.ws) {
        throw Error('not connected')
      }
      this.ws.send(JSON.stringify({
        action: action,
        data: data || {}
      }))
    },
    handle (ev) {
      const data = JSON.parse(ev.data)
      if (data.users) {
        this.users = data.users
      }
      if (data.result) {
        switch (data.result) {
          case 'connected':
            break
          case 'disconnected':
            this.disconnected()
        }
      }
    },
    disconnected () {
      this.ws.close()
      this.ws = undefined
      this.users = []
    },
    selectCard (value) {
      this.task.mark = value
    },
    getID () {
      let userID = localStorage.getItem('user_id')
      if (!userID) {
        userID = this.generateRandomString(16)
        localStorage.setItem('user_id', userID)
      }
      return userID
    },
    generateRandomString (length) {
      let result = ''
      let characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
      let charactersLength = characters.length
      for (let i = 0; i < length; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength))
      }
      return result
    }
  }
}
</script>

<style>
@import "bootstrap/dist/css/bootstrap.css";

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.poker-cards {
  display: flex;
  justify-content: space-between;
  align-content: center;
}

.poker-card {
  display: inline-block;
  width: 100px;
  height: 140px;
  font-weight: bold;
  font-size: 18px;
}
</style>
