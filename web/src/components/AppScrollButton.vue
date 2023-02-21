<template>
  <div class="btn scroll-top" :class="{ active: show }" @click="scrollUp()">
    <fa-icon icon="arrow-up" />
  </div>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted, ref } from "vue";

const show = ref(false);

onMounted(() => {
  window.addEventListener("scroll", handleScroll);
});

onUnmounted(() => {
  window.removeEventListener("scroll", handleScroll);
});

const handleScroll = () => {
  show.value = window.scrollY > 300;
};

const scrollUp = () => {
  window.scrollTo({ top: 0, behavior: "smooth" });
};
</script>

<style scoped>
.scroll-top {
  cursor: pointer;
  display: flex;
  position: fixed;
  bottom: -1em;
  right: 2em;

  background-color: rgb(74, 60, 92);

  padding: 0.5em;
  border-radius: 50%;

  visibility: hidden;
  opacity: 0;
  z-index: 9999;
  transition: 0.2s;
}

.scroll-top:hover {
  background-color: rgb(119, 70, 184);
}

.scroll-top.active {
  visibility: visible;
  opacity: 1;
  bottom: 2em;
}

.scroll-top svg {
  height: 36px;
  width: 36px;
}
</style>
