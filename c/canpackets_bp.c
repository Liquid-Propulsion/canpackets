// Code generated by bitproto. DO NOT EDIT.

#include "bitproto.h"
#include "canpackets_bp.h"

void BpXXXProcessArrayStagePacket1(void *data, struct BpProcessorContext *ctx) {
    struct BpArrayDescriptor descriptor = BpArrayDescriptor(false, 64, BpBool());
    BpEndecodeArray(&descriptor, ctx, data);
}

void BpXXXJsonFormatArrayStagePacket1(void *data, struct BpJsonFormatContext *ctx) {
    struct BpArrayDescriptor descriptor = BpArrayDescriptor(false, 64, BpBool());
    BpJsonFormatArray(&descriptor, ctx, data);
}

void BpFieldDescriptorsInitStagePacket(struct StagePacket *m, struct BpMessageFieldDescriptor *fds) {
    fds[0] = BpMessageFieldDescriptor((void *)&(m->solenoid_state), BpArray(64, 64 * sizeof(bool), BpXXXProcessArrayStagePacket1, BpXXXJsonFormatArrayStagePacket1), "solenoid_state");
}

void BpXXXProcessStagePacket(void *data, struct BpProcessorContext *ctx) {
    struct StagePacket *m = (struct StagePacket *)(data);
    struct BpMessageFieldDescriptor field_descriptors[1];
    BpFieldDescriptorsInitStagePacket(m, field_descriptors);
    struct BpMessageDescriptor descriptor = BpMessageDescriptor(false, 1, 64, field_descriptors);
    BpEndecodeMessage(&descriptor, ctx, data);
}

void BpXXXJsonFormatStagePacket(void *data, struct BpJsonFormatContext *ctx) {
    struct StagePacket *m = (struct StagePacket *)(data);
    struct BpMessageFieldDescriptor field_descriptors[1];
    BpFieldDescriptorsInitStagePacket(m, field_descriptors);
    struct BpMessageDescriptor descriptor = BpMessageDescriptor(false, 1, 64, field_descriptors);
    BpJsonFormatMessage(&descriptor, ctx, data);
}

int EncodeStagePacket(struct StagePacket *m, unsigned char *s) {
    struct BpProcessorContext ctx = BpProcessorContext(true, s);
    BpXXXProcessStagePacket((void *)m, &ctx);
    return 0;
}

int DecodeStagePacket(struct StagePacket *m, unsigned char *s) {
    struct BpProcessorContext ctx = BpProcessorContext(false, s);
    BpXXXProcessStagePacket((void *)m, &ctx);
    return 0;
}

int JsonStagePacket(struct StagePacket *m, char *s) {
    struct BpJsonFormatContext ctx = BpJsonFormatContext(s);
    BpXXXJsonFormatStagePacket((void *)m, &ctx);
    return ctx.n;
}

void BpFieldDescriptorsInitPowerPacket(struct PowerPacket *m, struct BpMessageFieldDescriptor *fds) {
    fds[0] = BpMessageFieldDescriptor((void *)&(m->system_powered), BpBool(), "system_powered");
    fds[1] = BpMessageFieldDescriptor((void *)&(m->siren), BpBool(), "siren");
}

void BpXXXProcessPowerPacket(void *data, struct BpProcessorContext *ctx) {
    struct PowerPacket *m = (struct PowerPacket *)(data);
    struct BpMessageFieldDescriptor field_descriptors[2];
    BpFieldDescriptorsInitPowerPacket(m, field_descriptors);
    struct BpMessageDescriptor descriptor = BpMessageDescriptor(false, 2, 2, field_descriptors);
    BpEndecodeMessage(&descriptor, ctx, data);
}

void BpXXXJsonFormatPowerPacket(void *data, struct BpJsonFormatContext *ctx) {
    struct PowerPacket *m = (struct PowerPacket *)(data);
    struct BpMessageFieldDescriptor field_descriptors[2];
    BpFieldDescriptorsInitPowerPacket(m, field_descriptors);
    struct BpMessageDescriptor descriptor = BpMessageDescriptor(false, 2, 2, field_descriptors);
    BpJsonFormatMessage(&descriptor, ctx, data);
}

int EncodePowerPacket(struct PowerPacket *m, unsigned char *s) {
    struct BpProcessorContext ctx = BpProcessorContext(true, s);
    BpXXXProcessPowerPacket((void *)m, &ctx);
    return 0;
}

int DecodePowerPacket(struct PowerPacket *m, unsigned char *s) {
    struct BpProcessorContext ctx = BpProcessorContext(false, s);
    BpXXXProcessPowerPacket((void *)m, &ctx);
    return 0;
}

int JsonPowerPacket(struct PowerPacket *m, char *s) {
    struct BpJsonFormatContext ctx = BpJsonFormatContext(s);
    BpXXXJsonFormatPowerPacket((void *)m, &ctx);
    return ctx.n;
}

void BpFieldDescriptorsInitBlinkPacket(struct BlinkPacket *m, struct BpMessageFieldDescriptor *fds) {
    fds[0] = BpMessageFieldDescriptor((void *)&(m->node_id), BpUint(8, sizeof(uint8_t)), "node_id");
}

void BpXXXProcessBlinkPacket(void *data, struct BpProcessorContext *ctx) {
    struct BlinkPacket *m = (struct BlinkPacket *)(data);
    struct BpMessageFieldDescriptor field_descriptors[1];
    BpFieldDescriptorsInitBlinkPacket(m, field_descriptors);
    struct BpMessageDescriptor descriptor = BpMessageDescriptor(false, 1, 8, field_descriptors);
    BpEndecodeMessage(&descriptor, ctx, data);
}

void BpXXXJsonFormatBlinkPacket(void *data, struct BpJsonFormatContext *ctx) {
    struct BlinkPacket *m = (struct BlinkPacket *)(data);
    struct BpMessageFieldDescriptor field_descriptors[1];
    BpFieldDescriptorsInitBlinkPacket(m, field_descriptors);
    struct BpMessageDescriptor descriptor = BpMessageDescriptor(false, 1, 8, field_descriptors);
    BpJsonFormatMessage(&descriptor, ctx, data);
}

int EncodeBlinkPacket(struct BlinkPacket *m, unsigned char *s) {
    struct BpProcessorContext ctx = BpProcessorContext(true, s);
    BpXXXProcessBlinkPacket((void *)m, &ctx);
    return 0;
}

int DecodeBlinkPacket(struct BlinkPacket *m, unsigned char *s) {
    struct BpProcessorContext ctx = BpProcessorContext(false, s);
    BpXXXProcessBlinkPacket((void *)m, &ctx);
    return 0;
}

int JsonBlinkPacket(struct BlinkPacket *m, char *s) {
    struct BpJsonFormatContext ctx = BpJsonFormatContext(s);
    BpXXXJsonFormatBlinkPacket((void *)m, &ctx);
    return ctx.n;
}

void BpFieldDescriptorsInitSensorDataPacket(struct SensorDataPacket *m, struct BpMessageFieldDescriptor *fds) {
    fds[0] = BpMessageFieldDescriptor((void *)&(m->node_id), BpUint(8, sizeof(uint8_t)), "node_id");
    fds[1] = BpMessageFieldDescriptor((void *)&(m->sensor_id), BpUint(4, sizeof(uint8_t)), "sensor_id");
    fds[2] = BpMessageFieldDescriptor((void *)&(m->sensor_data), BpUint(32, sizeof(uint32_t)), "sensor_data");
}

void BpXXXProcessSensorDataPacket(void *data, struct BpProcessorContext *ctx) {
    struct SensorDataPacket *m = (struct SensorDataPacket *)(data);
    struct BpMessageFieldDescriptor field_descriptors[3];
    BpFieldDescriptorsInitSensorDataPacket(m, field_descriptors);
    struct BpMessageDescriptor descriptor = BpMessageDescriptor(false, 3, 44, field_descriptors);
    BpEndecodeMessage(&descriptor, ctx, data);
}

void BpXXXJsonFormatSensorDataPacket(void *data, struct BpJsonFormatContext *ctx) {
    struct SensorDataPacket *m = (struct SensorDataPacket *)(data);
    struct BpMessageFieldDescriptor field_descriptors[3];
    BpFieldDescriptorsInitSensorDataPacket(m, field_descriptors);
    struct BpMessageDescriptor descriptor = BpMessageDescriptor(false, 3, 44, field_descriptors);
    BpJsonFormatMessage(&descriptor, ctx, data);
}

int EncodeSensorDataPacket(struct SensorDataPacket *m, unsigned char *s) {
    struct BpProcessorContext ctx = BpProcessorContext(true, s);
    BpXXXProcessSensorDataPacket((void *)m, &ctx);
    return 0;
}

int DecodeSensorDataPacket(struct SensorDataPacket *m, unsigned char *s) {
    struct BpProcessorContext ctx = BpProcessorContext(false, s);
    BpXXXProcessSensorDataPacket((void *)m, &ctx);
    return 0;
}

int JsonSensorDataPacket(struct SensorDataPacket *m, char *s) {
    struct BpJsonFormatContext ctx = BpJsonFormatContext(s);
    BpXXXJsonFormatSensorDataPacket((void *)m, &ctx);
    return ctx.n;
}

void BpFieldDescriptorsInitPongPacket(struct PongPacket *m, struct BpMessageFieldDescriptor *fds) {
    fds[0] = BpMessageFieldDescriptor((void *)&(m->node_id), BpUint(8, sizeof(uint8_t)), "node_id");
    fds[1] = BpMessageFieldDescriptor((void *)&(m->node_type), BpEnum(4, sizeof(NodeType)), "node_type");
}

void BpXXXProcessPongPacket(void *data, struct BpProcessorContext *ctx) {
    struct PongPacket *m = (struct PongPacket *)(data);
    struct BpMessageFieldDescriptor field_descriptors[2];
    BpFieldDescriptorsInitPongPacket(m, field_descriptors);
    struct BpMessageDescriptor descriptor = BpMessageDescriptor(false, 2, 12, field_descriptors);
    BpEndecodeMessage(&descriptor, ctx, data);
}

void BpXXXJsonFormatPongPacket(void *data, struct BpJsonFormatContext *ctx) {
    struct PongPacket *m = (struct PongPacket *)(data);
    struct BpMessageFieldDescriptor field_descriptors[2];
    BpFieldDescriptorsInitPongPacket(m, field_descriptors);
    struct BpMessageDescriptor descriptor = BpMessageDescriptor(false, 2, 12, field_descriptors);
    BpJsonFormatMessage(&descriptor, ctx, data);
}

int EncodePongPacket(struct PongPacket *m, unsigned char *s) {
    struct BpProcessorContext ctx = BpProcessorContext(true, s);
    BpXXXProcessPongPacket((void *)m, &ctx);
    return 0;
}

int DecodePongPacket(struct PongPacket *m, unsigned char *s) {
    struct BpProcessorContext ctx = BpProcessorContext(false, s);
    BpXXXProcessPongPacket((void *)m, &ctx);
    return 0;
}

int JsonPongPacket(struct PongPacket *m, char *s) {
    struct BpJsonFormatContext ctx = BpJsonFormatContext(s);
    BpXXXJsonFormatPongPacket((void *)m, &ctx);
    return ctx.n;
}