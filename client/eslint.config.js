//  @ts-check

import { tanstackConfig } from '@tanstack/eslint-config'
import { defineConfig, globalIgnores } from "eslint/config";

const customConfig = defineConfig([globalIgnores(["src/proto/", "src/generated/"])])

export default [...tanstackConfig, ...customConfig]
