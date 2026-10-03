import { useEffect, useMemo, useState } from 'react';
import Box from '@mui/material/Box';
import FormHelperText from '@mui/material/FormHelperText';
import Typography from '@mui/material/Typography';
import { Form } from '@rjsf/mui';
import validator from '@rjsf/validator-ajv8';
import { RJSFSchema } from '@rjsf/utils';
import { InputHelperText, required, useInput } from 'react-admin';
import { useFormContext, useWatch } from 'react-hook-form';
import { apiUrl, httpClient } from '../dataProvider';
import { AdvancedJsonEditor } from './AdvancedJsonEditor';
import { EDITOR_ERRORS_FIELD, StringMap, validateStringMap } from './StringMapInput';

type ExecutorPlugin = {
    id: string;
    name: string;
    schema?: RJSFSchema;
};

const parseSchemaValue = (value: unknown, propertySchema: RJSFSchema): unknown => {
    if (typeof value !== 'string') {
        return value;
    }

    const type = Array.isArray(propertySchema.type) ? propertySchema.type[0] : propertySchema.type;
    if (type === 'boolean') {
        if (value === 'true') return true;
        if (value === 'false') return false;
    }
    if (type === 'number' || type === 'integer') {
        const parsed = Number(value);
        return Number.isNaN(parsed) ? value : parsed;
    }
    if (type === 'array' || type === 'object') {
        try {
            return JSON.parse(value);
        } catch {
            return value;
        }
    }
    return value;
};

const toSchemaFormData = (value: StringMap | undefined, schema: RJSFSchema) => {
    const properties = (schema.properties || {}) as Record<string, RJSFSchema>;
    return Object.entries(value || {}).reduce<Record<string, unknown>>((acc, [key, currentValue]) => {
        acc[key] = parseSchemaValue(currentValue, properties[key] || {});
        return acc;
    }, {});
};

const toExecutorConfig = (formData: Record<string, unknown> | undefined) =>
    Object.entries(formData || {}).reduce<StringMap>((acc, [key, value]) => {
        if (value === undefined || value === null) {
            return acc;
        }
        if (typeof value === 'string') {
            acc[key] = value;
            return acc;
        }
        if (typeof value === 'boolean' || typeof value === 'number') {
            acc[key] = String(value);
            return acc;
        }
        acc[key] = JSON.stringify(value);
        return acc;
    }, {});

const JsonExecutorConfigInput = () => {
    const {
        field,
        fieldState: { error },
        isRequired,
    } = useInput({ source: 'executor_config', validate: required(), defaultValue: {} });
    const { clearErrors, setError, setValue, trigger } = useFormContext();

    const handleValidityChange = (valid: boolean, message?: string) => {
        const validationMessage = valid ? undefined : (message || 'Fix the invalid executor configuration JSON.');
        setValue(`${EDITOR_ERRORS_FIELD}.executor_config`, validationMessage, {
            shouldDirty: false,
            shouldTouch: false,
            shouldValidate: false,
        });
        if (validationMessage) {
            setError('executor_config', { type: 'validate', message: validationMessage });
        } else {
            clearErrors('executor_config');
        }
        void trigger('executor_config');
    };

    return (
        <Box sx={{ width: '100%' }}>
            <Typography variant="subtitle2" sx={{ mb: 1 }}>
                Executor configuration{isRequired ? ' *' : ''}
            </Typography>
            <AdvancedJsonEditor<StringMap>
                label="Edit executor config as JSON"
                value={(field.value ?? {}) as StringMap}
                validateValue={validateStringMap}
                onChange={field.onChange}
                onValidityChange={handleValidityChange}
            />
            <FormHelperText error={Boolean(error)}>
                <InputHelperText
                    error={error?.message}
                    helperText="Executor is not registered or has no schema. Edit its string map configuration as JSON."
                />
            </FormHelperText>
        </Box>
    );
};

const SchemaExecutorConfigInput = ({ schema }: { schema: RJSFSchema }) => {
    const { field, fieldState } = useInput({
        source: 'executor_config',
        validate: required(),
        defaultValue: {},
    });
    const formData = useMemo(
        () => toSchemaFormData(field.value, schema),
        [field.value, schema]
    );

    return (
        <Box sx={{ width: '100%' }}>
            <Typography variant="subtitle2" sx={{ mb: 1 }}>
                Executor configuration
            </Typography>
            <Form
                schema={schema}
                validator={validator}
                formData={formData}
                liveValidate
                noHtml5Validate
                showErrorList={false}
                onBlur={() => field.onBlur()}
                onChange={event => field.onChange(toExecutorConfig(event.formData))}
            >
                <></>
            </Form>
            <FormHelperText error={!!fieldState.error}>
                {fieldState.error?.message || 'Configuration arguments for the selected executor.'}
            </FormHelperText>
        </Box>
    );
};

export const ExecutorConfigInput = () => {
    const executor = useWatch({ name: 'executor' });
    const [plugins, setPlugins] = useState<ExecutorPlugin[]>([]);

    useEffect(() => {
        httpClient(`${apiUrl}/plugins/executors`)
            .then(({ json }) => setPlugins(json))
            .catch(() => setPlugins([]));
    }, []);

    const selectedPlugin = plugins.find(plugin => plugin.name === executor || plugin.id === executor);
    const schema = selectedPlugin?.schema;

    if (!executor || !schema) {
        return <JsonExecutorConfigInput />;
    }

    return <SchemaExecutorConfigInput schema={schema} />;
};
