--
-- PostgreSQL database dump
--

-- Dumped from database version 16rc1
-- Dumped by pg_dump version 16rc1

-- Started on 2024-11-10 01:30:41

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 4852 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 216 (class 1259 OID 26074)
-- Name: Category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Category" (
    id integer NOT NULL,
    name character varying NOT NULL,
    description character varying NOT NULL
);


ALTER TABLE public."Category" OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 26073)
-- Name: Category_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Category_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."Category_id_seq" OWNER TO postgres;

--
-- TOC entry 4853 (class 0 OID 0)
-- Dependencies: 215
-- Name: Category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Category_id_seq" OWNED BY public."Category".id;


--
-- TOC entry 218 (class 1259 OID 26083)
-- Name: Goods; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Goods" (
    id integer NOT NULL,
    category_id integer NOT NULL,
    name character varying NOT NULL,
    photo_url character varying NOT NULL,
    price character varying NOT NULL,
    purchase_date date NOT NULL,
    total_usage_days integer NOT NULL
);


ALTER TABLE public."Goods" OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 26082)
-- Name: Goods_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Goods_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."Goods_id_seq" OWNER TO postgres;

--
-- TOC entry 4854 (class 0 OID 0)
-- Dependencies: 217
-- Name: Goods_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Goods_id_seq" OWNED BY public."Goods".id;


--
-- TOC entry 4693 (class 2604 OID 26077)
-- Name: Category id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Category" ALTER COLUMN id SET DEFAULT nextval('public."Category_id_seq"'::regclass);


--
-- TOC entry 4694 (class 2604 OID 26086)
-- Name: Goods id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Goods" ALTER COLUMN id SET DEFAULT nextval('public."Goods_id_seq"'::regclass);


--
-- TOC entry 4844 (class 0 OID 26074)
-- Dependencies: 216
-- Data for Name: Category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Category" (id, name, description) FROM stdin;
1	Furniture	Furniture for office
2	Electronics	Office electronics
3	Appliances	Appliances for office
\.


--
-- TOC entry 4846 (class 0 OID 26083)
-- Dependencies: 218
-- Data for Name: Goods; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Goods" (id, category_id, name, photo_url, price, purchase_date, total_usage_days) FROM stdin;
1	1	Office Desk	http://example.com/office_desk.jpg	2500000	2024-09-25	45
2	1	Office Chair	http://example.com/office_chair.jpg	1500000	2024-08-11	90
3	2	Desktop Computer	http://example.com/desktop_computer.jpg	8000000	2024-06-11	120
4	3	Coffee Machine	http://example.com/coffee_machine.jpg	1800000	2024-10-10	30
5	2	Laptop	http://example.com/laptop_dell_xps_13.jpg	12000000	2024-08-26	75
6	2	Monitor 24 Inch	http://example.com/monitor_24inch.jpg	3500000	2024-09-25	45
7	3	Refrigerator	http://example.com/office_fridge.jpg	5000000	2024-06-30	100
\.


--
-- TOC entry 4855 (class 0 OID 0)
-- Dependencies: 215
-- Name: Category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Category_id_seq"', 3, true);


--
-- TOC entry 4856 (class 0 OID 0)
-- Dependencies: 217
-- Name: Goods_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Goods_id_seq"', 7, true);


--
-- TOC entry 4696 (class 2606 OID 26081)
-- Name: Category Category_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Category"
    ADD CONSTRAINT "Category_pkey" PRIMARY KEY (id);


--
-- TOC entry 4698 (class 2606 OID 26090)
-- Name: Goods Goods_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Goods"
    ADD CONSTRAINT "Goods_pkey" PRIMARY KEY (id);


--
-- TOC entry 4699 (class 2606 OID 26091)
-- Name: Goods Goods_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Goods"
    ADD CONSTRAINT "Goods_category_id_fkey" FOREIGN KEY (category_id) REFERENCES public."Category"(id) ON DELETE CASCADE;


-- Completed on 2024-11-10 01:30:42

--
-- PostgreSQL database dump complete
--

