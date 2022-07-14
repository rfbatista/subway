import dotenv from "dotenv";

const getEnv = (envVar: string, default = null) => {
  return ''
}

const config = {
  host: getEnv('HOST'),
  port: getEnv('PORT'),
};
export default config;
