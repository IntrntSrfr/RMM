<template>
  <div class="member" :class="{ selected: checked }">
    <div class="icon">
      <img :src="iconUrl" alt="avatar" class="icon-inner" />
    </div>
    <div class="body">
      <div class="name">
        {{ username }}
      </div>
      <div class="joined">Joined {{ timeAgo.format(joined) }}</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, defineProps } from "vue";
import TimeAgo from "javascript-time-ago";

const timeAgo = new TimeAgo("en-US");

const props = defineProps<{
  checked: boolean;
  id: string;
  username: string;
  icon: string;
  joined: Date;
}>();

const iconUrl = computed(() => {
  if (!props.icon) {
    return `https://cdn.discordapp.com/embed/avatars/${
      parseInt(props.id) % 5
    }.png?size=64`;
  }
  return props.icon.startsWith("a_")
    ? `https://cdn.discordapp.com/avatars/${props.id}/${props.icon}.gif?size=64`
    : `https://cdn.discordapp.com/avatars/${props.id}/${props.icon}.webp?size=64`;
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
  overflow: hidden;
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
  flex-shrink: 0;
}

.icon-inner {
  height: 100%;
  width: 100%;
  border-radius: 50%;
}

.body {
  margin-left: 0.5em;
  overflow: hidden;
  flex: 1 0 auto;
}

.name {
  font-size: 1rem;
  color: rgb(227, 227, 227);
  text-overflow: ellipsis;
  overflow: hidden;
}

.joined {
  font-size: 0.75rem;
  color: rgb(174, 174, 174);
}
</style>
