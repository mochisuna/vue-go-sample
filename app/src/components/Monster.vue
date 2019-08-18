<template>
  <div class="monster">
    <h1>{{ msg }}</h1>
    <h2>名前</h2>
    <h3>攻撃力: {{attack}}</h3>
    <input v-model="monster_name" />
    <button v-on:click="doAdd">モンスターを追加</button>
    <ul>
      <li
        v-for="(item,index) in monster_list"
        v-bind:key="item.id"
        v-bind:class="{ 'strong-monster': item.max > 250, 'dying': item.hp/item.max <= 0.2 }"
      >
        <div v-if="!item.dead">
          ID.{{ item.id }} {{ item.name }} [ HP.{{ item.hp }} / {{ item.max }} ]
          <span
            v-if="item.hp/item.max <= 0.2"
            dying
          >瀕死！</span>
          <button v-on:click="doAttack(index)">攻撃する</button>
          <button v-on:click="doRemove(index)">削除</button>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "Monster",
  props: {
    msg: String
  },
  data() {
    return {
      attack: 10,
      monster_name: "キマイラ",
      monster_list: []
    };
  },
  methods: {
    // 追加ボタンをクリックしたときのハンドラ
    doAdd: function() {
      console.log("appeared");
      // リスト内で1番大きいIDを取得
      var last = this.monster_list.reduce(function(a, b) {
        return a > b.id ? a : b.id;
      }, 0);
      // 新しいモンスターをリストに追加
      if (this.monster_name) {
        var max = Math.floor(Math.random() * Math.floor(500)) + 1;
        var hp = Math.floor(Math.random() * Math.floor(max)) + 1;
        this.monster_list.push({
          id: last + 1, // 現在の最大のIDに+1してユニークなIDを作成
          name: this.monster_name, // 現在のフォームの入力値
          hp: hp,
          max: max,
          dead: false
        });
      }
    },
    doRemove: function(index) {
      // 受け取ったインデックスの位置から1個要素を削除
      this.monster_list.splice(index, 1);
    },
    doAttack: function(index) {
      this.monster_list[index].hp -= this.attack; // HPを減らす
      if (this.monster_list[index].hp <= 0) {
        console.log("Lv up!!!");
        this.monster_list[index].dead = true;
        this.attack++;
      }
    }
  },
  created: function() {
    this.axios
      .get(process.env.VUE_APP_API_ORIGIN)
      .then(
        function(response) {
          // 取得完了したらlistリストに代入
          console.log(response);
          this.monster_list = response.data;
        }.bind(this)
      )
      .catch(function(e) {
        console.error(e);
      });
  }
};
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
a {
  color: #42b983;
}
.strong-monster {
  background: #ffffcc;
  font-weight: bold;
}

.dying {
  color: #ff7777;
}

.v-enter-active,
.v-leave-active {
  transition: opacity 1s;
}

.v-enter,
.v-leave-to {
  opacity: 0;
}
</style>
