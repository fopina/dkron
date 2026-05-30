import { useEffect, useMemo, useState } from "react";
import Box from "@mui/material/Box";
import FormHelperText from "@mui/material/FormHelperText";
import Typography from "@mui/material/Typography";
import Form from "@rjsf/mui";
import validator from "@rjsf/validator-ajv8";
import { RJSFSchema } from "@rjsf/utils";
import { required, useInput } from "react-admin";
import { JsonInput } from "react-admin-json-view";
import { useWatch } from "react-hook-form";
import { apiUrl, httpClient } from "../dataProvider";

type ExecutorPlugin = {
  id: string;
  name: string;
  schema?: RJSFSchema;
};

const parseSchemaValue = (
  value: unknown,
  propertySchema: RJSFSchema
): unknown => {
  if (typeof value !== "string") {
    return value;
  }

  const type = Array.isArray(propertySchema.type)
    ? propertySchema.type[0]
    : propertySchema.type;
  if (type === "boolean") {
    if (value === "true") return true;
    if (value === "false") return false;
  }
  if (type === "number" || type === "integer") {
    const parsed = Number(value);
    return Number.isNaN(parsed) ? value : parsed;
  }
  if (type === "array" || type === "object") {
    try {
      return JSON.parse(value);
    } catch {
      return value;
    }
  }
  return value;
};

const toSchemaFormData = (
  value: Record<string, unknown> | undefined,
  schema: RJSFSchema
) => {
  const properties = (schema.properties || {}) as Record<string, RJSFSchema>;
  return Object.entries(value || {}).reduce<Record<string, unknown>>(
    (acc, [key, currentValue]) => {
      acc[key] = parseSchemaValue(currentValue, properties[key] || {});
      return acc;
    },
    {}
  );
};

const toExecutorConfig = (formData: Record<string, unknown> | undefined) =>
  Object.entries(formData || {}).reduce<Record<string, string>>(
    (acc, [key, value]) => {
      if (value === undefined || value === null) {
        return acc;
      }
      if (typeof value === "string") {
        acc[key] = value;
        return acc;
      }
      if (typeof value === "boolean" || typeof value === "number") {
        acc[key] = String(value);
        return acc;
      }
      acc[key] = JSON.stringify(value);
      return acc;
    },
    {}
  );

const JsonExecutorConfigInput = () => (
  <JsonInput
    source="executor_config"
    reactJsonOptions={{
      name: null,
      collapsed: true,
      enableClipboard: false,
      displayDataTypes: false,
    }}
    helperText="Configuration arguments for the specific executor."
    validate={required()}
  />
);

const SchemaExecutorConfigInput = ({ schema }: { schema: RJSFSchema }) => {
  const { field, fieldState } = useInput({
    source: "executor_config",
    validate: required(),
  });
  const formData = useMemo(
    () => toSchemaFormData(field.value, schema),
    [field.value, schema]
  );

  return (
    <Box sx={{ mb: 2, mt: 1 }}>
      <Typography
        component="label"
        variant="body2"
        sx={{ color: "text.secondary" }}
      >
        Executor Config
      </Typography>
      <Form
        schema={schema}
        validator={validator}
        formData={formData}
        liveValidate
        noHtml5Validate
        showErrorList={false}
        onBlur={() => field.onBlur()}
        onChange={(event) => field.onChange(toExecutorConfig(event.formData))}
      >
        <></>
      </Form>
      <FormHelperText error={!!fieldState.error}>
        {fieldState.error?.message ||
          "Configuration arguments for the selected executor."}
      </FormHelperText>
    </Box>
  );
};

export const ExecutorConfigInput = () => {
  const executor = useWatch({ name: "executor" });
  const [plugins, setPlugins] = useState<ExecutorPlugin[]>([]);

  useEffect(() => {
    httpClient(`${apiUrl}/plugins/executors`)
      .then(({ json }) => setPlugins(json))
      .catch(() => setPlugins([]));
  }, []);

  const selectedPlugin = plugins.find(
    (plugin) => plugin.name === executor || plugin.id === executor
  );
  const schema = selectedPlugin?.schema;

  if (!executor || !schema) {
    return <JsonExecutorConfigInput />;
  }

  return <SchemaExecutorConfigInput schema={schema} />;
};
