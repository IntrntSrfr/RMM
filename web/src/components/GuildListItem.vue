<template>
  <div class="guild">
    <div class="icon">
      <img class="icon-inner" :src="guildIcon" alt="Server icon" />
    </div>
    <div class="name">
      {{ name }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, defineProps, withDefaults } from "vue";

export interface Props {
  id: string;
  name: string;
  icon?: string; // { type: String, default: "" },
}

const props = withDefaults(defineProps<Props>(), { icon: "" });

const guildIcon = computed(() => {
  if (!props.icon) {
    return `https://cdn.discordapp.com/embed/avatars/${
      parseInt(props.id) % 5
    }.png?size=128`;
  }
  return props.icon.startsWith("a_")
    ? `https://cdn.discordapp.com/icons/${props.id}/${props.icon}.gif?size=128`
    : `https://cdn.discordapp.com/icons/${props.id}/${props.icon}.webp?size=128`;
});
</script>

<style scoped>
.guild {
  display: flex;
  align-items: center;
  padding: 0.5em;
  transition: 0.1s;
}

.guild:hover {
  background-color: rgba(30, 143, 255, 0.1);
}

.icon {
  height: 48px;
  width: 48px;
}

.icon-inner {
  height: 100%;
  width: 100%;
  border-radius: 50%;
}

.name {
  font-size: 1.125rem;
  margin-left: 0.5em;
}

@media only screen and (min-width: 660px) {
  .guild {
    border: 1px solid dodgerblue;
    border-radius: 5px;
  }
}
</style>
