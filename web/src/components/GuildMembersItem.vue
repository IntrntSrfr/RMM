<template>
  <div class="member" :class="{ selected: checked }">
    <div class="icon">
      <img :src="iconUrl" alt="avatar" class="icon-inner" />
    </div>
    <div class="body">
      <div class="name">
        {{ username }}
      </div>
      <div class="joined">{{ timeAgo.format(joined) }}</div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import TimeAgo from "javascript-time-ago";

export default defineComponent({
  name: "GuildMembersItem",
  setup() {
    const timeAgo = new TimeAgo("en-US");
    return { timeAgo };
  },
  props: {
    checked: Boolean,
    id: { type: String, required: true },
    username: String,
    icon: String,
    joined: { type: Date, required: true },
  },
  computed: {
    iconUrl(): string {
      if (!this.icon) {
        return `https://cdn.discordapp.com/embed/avatars/${
          parseInt(this.id) % 5
        }.png?size=64`;
      }
      return this.icon.startsWith("a_")
        ? `https://cdn.discordapp.com/avatars/${this.id}/${this.icon}.gif?size=64`
        : `https://cdn.discordapp.com/avatars/${this.id}/${this.icon}.webp?size=64`;
    },
  },
});
</script>

<style scoped>
.member {
  display: flex;
  align-items: center;
  cursor: pointer;
  padding: 0.75em;
  transition: 0.1s;
  border-color: rgba(30, 143, 255, 0.2);
}

.member:hover {
  background-color: rgba(30, 143, 255, 0.1);
}

.member.selected {
  border-left: 5px solid rgba(30, 143, 255, 0.2);
}

.icon {
  height: 36px;
  width: 36px;
}

.icon-inner {
  height: 100%;
  width: 100%;
  border-radius: 50%;
}

.body {
  margin-left: 0.5em;
}

.name {
  font-size: 1rem;
  color: rgb(227, 227, 227);
}
.joined {
  font-size: 0.75rem;
  color: rgb(174, 174, 174);
}
</style>
