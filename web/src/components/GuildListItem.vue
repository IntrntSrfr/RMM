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

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "GuildListItem",
  props: {
    id: { type: String, required: true },
    name: { type: String, required: true },
    icon: { type: String, default: "" },
  },
  computed: {
    guildIcon(): string {
      if (!this.icon) {
        return `https://cdn.discordapp.com/embed/avatars/${
          parseInt(this.id) % 5
        }.png`;
      }
      return this.icon.startsWith("a_")
        ? `https://cdn.discordapp.com/icons/${this.id}/${this.icon}.gif`
        : `https://cdn.discordapp.com/icons/${this.id}/${this.icon}.png`;
    },
  },
});
</script>

<style scoped>
.guild {
  display: flex;
  align-items: center;
  padding: 0.5em;
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
