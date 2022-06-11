<template>
  <div class="member" :class="{ selected: checked }">
    <div class="name">
      {{ username }}
    </div>
    <div class="joined">{{ timeAgo.format(joined) }}</div>
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
    username: String,
    joined: { type: Date, required: true },
  },
});
</script>

<style scoped>

.member {
  cursor: pointer;
  padding: 0.75em;
  transition: 0.1s;
}

.member:hover {
  background-color: rgba(30, 143, 255, 0.1);
}

.member.selected {
  background-color: rgba(30, 143, 255, 0.2);
}

.name {
  font-size: 1rem;
  color: rgb(175, 175, 175);
}
.joined {
  font-size: 0.75rem;
  color: #666;
}
</style>
