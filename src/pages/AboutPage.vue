<script setup lang="ts">
import { getToken, mStorage } from '@/utils/tools';
import { outLogin } from '@/api/Account';
const token = getToken();
const UserInfo = mStorage.get('user_info');

const handleOutLogin = () => {
  outLogin();
};
</script>

<template>
  <div class="about">
    <div v-if="token">
      <br />
      <img class="Avatar" :src="UserInfo.Avatar" alt="" srcset="" />
      <br />
      欢迎访问,{{ UserInfo.NickName }}
      <br />
      <button @click="handleOutLogin">退出登录</button>
    </div>
    <div v-if="!token">
      <br />
      <a href="/login"><button>请登录</button></a>
      <br />
    </div>
  </div>
</template>

<style lang="less">
.Avatar {
  width: 100px;
  height: 100px;
  border-radius: 100px;
}
</style>
