const user = {
  mail: "e850506@gmail.com",
  name: "Eric",

  login() {
    console.log(`${this.name} is logged in with email: ${this.mail}.`);
  },

  logout() {
    console.log(`${this.name} is logged out with email: ${this.mail}.`);
  },
};

console.log(user);
user.login();
user.logout();
