<template>
  <div class="member-wrapper">
    <div class="options">
      <div class="grp">
        <AppButton text="Select all" @click="markAll" />
        <AppButton text="Select none" @click="markNone" />
      </div>
      <div class="grp">
        <AppButton text="Ban selected" variant="danger" @click="banSelected" />
      </div>
    </div>
    <div class="member-list">
      <GuildMembersItem
        v-for="(m, i) in members"
        :key="i"
        @click="mark(m)"
        :checked="m.checked"
        :username="`${m.member.user.username}#${m.member.user.discriminator}`"
        :joined="m.joined"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import type { Member } from "@/stores/user";
import AppButton from "@/components/AppButton.vue";
import { useGuildStore } from "@/stores/guilds";
import GuildMembersItem from "./GuildMembersItem.vue";

type TableMember = {
  member: Member;
  checked: boolean;
  joined: Date;
};

export default defineComponent({
  name: "GuildMembers",
  components: { AppButton, GuildMembersItem },
  setup() {
    const guildStore = useGuildStore();
    return { guildStore };
  },
  data() {
    return {
      members: [] as TableMember[],
    };
  },
  computed: {
    selectedMembers(): TableMember[] {
      return this.members.filter((m: TableMember) => m.checked);
    },
  },
  methods: {
    markAll() {
      this.members.forEach((m: TableMember) => (m.checked = true));
    },
    markNone(): void {
      this.members.forEach((m: TableMember) => (m.checked = false));
    },
    mark(tm: TableMember){
      tm.checked = !tm.checked;
    },
    banSelected() {
      if (!this.selectedMembers.length) {
        return;
      }
      const ans = confirm("Are you SURE you want to ban ALL selected users?");
      if (!ans) {
        return;
      }

      console.log(this.selectedMembers);
      console.log(this.selectedMembers.map((m) => m.member.user.id).join(" "));
    },
  },
  async created() {
    const guildID = this.$route.params.guildID;
    await this.guildStore.fetchGuildMembers(guildID.toString());
    const g = this.guildStore.getGuildByID(guildID.toString());

    if (!g || !g.members) {
      return;
    }

    this.members = g.members
      .map((m: Member) => {
        return {
          member: m,
          checked: false,
          joined: new Date(m.joined_at),
        } as TableMember;
      })
      .sort((a: TableMember, b: TableMember) => {
        return b.joined.getTime() - a.joined.getTime();
      });
  },
});
</script>

<style scoped>
.options {
  display: flex;
  justify-content: space-between;
  gap: 1em;
  padding-bottom: 1em;
  border-bottom: 1px solid #666;
}

.grp {
  display:flex;
  gap:1em;
}

.member-list {
  margin-top: 0.5em;
  display: flex;
  flex-direction: column;
  width: 100%;
}

.member-list .member:first-child {
  border-radius: 10px 10px 0 0;
}

.member-list .member:last-child {
  border-radius: 0 0 10px 10px;
}

.member-list .member:only-child {
  border-radius: 10px;
}

.member + .member {
  border-top: 1px solid dodgerblue;
}

</style>
