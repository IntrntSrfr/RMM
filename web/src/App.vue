<template>
  <RouterView />
</template>

<style>
@import "@/assets/base.css";

#app {
  max-width: 900px;
  margin: 0 auto;
  padding: 0.5rem 1rem;

  font-weight: normal;
}

@media only screen and (min-width: 900px) {
  #app {
    padding: 2rem;
  }
}
</style>
