CREATE TABLE table_company
(
    company varchar(36),

    CONSTRAINT primery_key_company PRIMARY KEY (company)
);

CREATE TABLE table_cpu_plus_specifications
(
    id_components int, -- seq_id_components

    CONSTRAINT primery_key_id_components PRIMARY KEY (id_components)
);

CREATE TABLE table_cpu_plus_specifications
(
    id_cpu int GENERATED ALWAYS AS IDENTITY,
    company varchar(128),
    lineup varchar(128),
    processor_generation smallint,
    processor_sku smallint,
    processor_suffix varchar(4);
    socket varchar(7),
    cores smallint,
    logical_processors smallint,
    base_speed numeric,
    max_speed numeric,
    l2_cache smallint,
    l3_cache smallint,
    support_ddr4 bool,
    support_ddr5 bool,
    tdp smallint,
    nm_process smallint,
    market_launch_date smallint
);

CREATE TABLE table_motherboard_plus_specifications
(
    id_motherboard int GENERATED ALWAYS AS IDENTITY,
    company varchar(128),
    lineup varchar(128),
    form_factor varchar(128),
    chipset varchar(128),
    max_speed_memory smallint,
    memory_slots smallint,
    pci_express_version smallint,
    m2_slots smallint,
    market_launch_date smallint
);

CREATE TABLE table_rum_plus_specifications
(
    id_motherboard int GENERATED ALWAYS AS IDENTITY,
    company varchar(128),
    lineup varchar(128),
    ddr4 bool,
    ddr5 bool,
    base_speed smallint,
    pc_index varchar(9),
    cas_latency varchar(3),
    timings varchar(8),
    supply_voltage numeric,

    CONSTRAINT check_ddr4_ddr5 CHECK (ddr4 != ddr5)
);