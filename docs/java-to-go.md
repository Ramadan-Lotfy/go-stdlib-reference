# Java to Go Cheat Sheet

| Java | Go |
|------|----|
| `String.format`, `System.out.printf` | `fmt.Sprintf`, `fmt.Printf` |
| `StringBuilder` | `strings.Builder` |
| `String.split(regex)` | `strings.Split` (literal separator), `regexp` for patterns |
| `String.strip()` | `strings.TrimSpace` |
| `String.repeat(n)` | `strings.Repeat` |
| `equalsIgnoreCase` | `strings.EqualFold` |
| `Integer.parseInt` | `strconv.Atoi` (returns an error) |
| `Integer.toString(n, radix)` | `strconv.FormatInt(n, base)` |
| `Double.parseDouble` | `strconv.ParseFloat(s, 64)` |
| `Boolean.parseBoolean` | `strconv.ParseBool` (unknown input is an error) |
| `NumberFormatException` | `*strconv.NumError` (`errors.Is(err, strconv.ErrSyntax)`) |
| `try/catch` | `if err != nil { ... }` |
| `java.util.List/Map` | slices and maps (`slices`, `maps`) |
| `java.io.InputStream/OutputStream` | `io.Reader` / `io.Writer` |
| `ByteArrayOutputStream` | `bytes.Buffer` |
| `ByteArrayInputStream` | `bytes.NewReader` |
| `Arrays.equals(byte[], byte[])` | `bytes.Equal` |
| `s.getBytes(UTF_8)` / `new String(b, UTF_8)` | `[]byte(s)` / `string(b)` |
| `ExecutorService`, `Future` | goroutines, channels, `sync.WaitGroup` |
| `CompletableFuture` cancellation | `context.Context` |
| Jackson | `encoding/json` |
| SLF4J | `log/slog` |
| Spring MVC / Tomcat | `net/http` (`ServeMux`, `Handler`) |
| JDBC / HikariCP | `database/sql` (built-in pool) |
| JUnit | `testing` |

---

## By package

Each package README has its own Java comparison; here they are in one place.

### errors

| Java | Go |
|---|---|
| `try { } catch (IOException e)` | `if err != nil { }` |
| `throw new X()` | `return X{}` (an error value) |
| `catch (NotFound e)` by type | `errors.As(err, &target)` |
| `e.getCause()` | `errors.Unwrap(err)` |
| `addSuppressed` | `errors.Join` |

### os

| Java | Go |
|---|---|
| `Files.readAllBytes`, `Files.writeString` | `os.ReadFile`, `os.WriteFile` |
| `Files.createTempFile` | `os.CreateTemp` |
| `System.getenv("X")` | `os.Getenv("X")` |
| `File.exists()` | `os.Stat` and `errors.Is(err, fs.ErrNotExist)` |
| `System.exit(1)` | `os.Exit(1)` |
| `args` in `main(String[] args)` | `os.Args[1:]` |

### sort

| Java | Go |
|---|---|
| `Collections.sort(list)` | `slices.Sort(s)` or `sort.Ints(s)` |
| `list.sort(Comparator)` | `sort.Slice(s, less)` or `slices.SortFunc` |
| `Comparator.reversed()` | `sort.Sort(sort.Reverse(...))` |
| `Collections.binarySearch` | `sort.SearchInts` / `slices.BinarySearch` |

### slices

| Java | Go |
|---|---|
| `Collections.sort` | `slices.Sort` |
| `list.contains(x)` | `slices.Contains(s, x)` |
| `list.indexOf(x)` | `slices.Index(s, x)` |
| `Collections.max` | `slices.Max(s)` |
| `Collections.reverse` | `slices.Reverse(s)` |
| `list.removeIf` | `slices.DeleteFunc` |

### maps

| Java | Go |
|---|---|
| `HashMap<K,V>` | `map[K]V` |
| `map.getOrDefault(k, d)` | `v, ok := m[k]` then default if `!ok` |
| `map.containsKey(k)` | `_, ok := m[k]` |
| `new HashMap<>(other)` | `maps.Clone(other)` |
| `HashSet<T>` | `map[T]struct{}` |
| `TreeMap` (sorted) | map + sorted key slice, or a third-party tree |

### container/list, container/heap, container/ring

| Java | Go |
|---|---|
| `PriorityQueue<T>` | `container/heap` on a slice type |
| `LinkedList<T>` | `container/list` |
| `LinkedHashMap` as LRU | `list` + `map` (see the exercise) |

### time

| Java | Go |
|---|---|
| `Instant.now()`, `LocalDateTime` | `time.Now()` |
| `Duration.ofMinutes(90)` | `90 * time.Minute` |
| `DateTimeFormatter.ofPattern("yyyy-MM-dd")` | `t.Format("2006-01-02")` |
| `ZoneId`, `ZonedDateTime` | `*time.Location`, `t.In(loc)` |
| `ScheduledExecutorService` | `time.Ticker` / `time.AfterFunc` |

### io

| Java | Go |
|---|---|
| `InputStream` / `OutputStream` | `io.Reader` / `io.Writer` |
| `InputStream.transferTo(out)` | `io.Copy(dst, src)` |
| `readAllBytes()` | `io.ReadAll(r)` |
| `BoundedInputStream` | `io.LimitReader` |
| `PipedInputStream`/`PipedOutputStream` | `io.Pipe` |
| `OutputStream.nullOutputStream()` | `io.Discard` |

### bufio

| Java | Go |
|---|---|
| `BufferedReader` | `bufio.Reader` / `bufio.Scanner` |
| `reader.readLine()` | `scanner.Scan()` + `scanner.Text()` |
| `java.util.Scanner` | `bufio.Scanner` (tokens only, no number parsing) |
| `BufferedWriter.flush()` | `bufio.Writer.Flush()` |

### io/fs

| Java | Go |
|---|---|
| `java.nio.file.FileSystem` | `fs.FS` |
| `Files.walk` | `fs.WalkDir` |
| `Class.getResourceAsStream` | `embed.FS` + `fs.ReadFile` |
| `Jimfs` (in-memory FS for tests) | `fstest.MapFS` |

### path/filepath

| Java | Go |
|---|---|
| `Paths.get(a, b)`, `Path.resolve` | `filepath.Join(a, b)` |
| `path.getFileName()` | `filepath.Base(p)` |
| `path.getParent()` | `filepath.Dir(p)` |
| `path.relativize` | `filepath.Rel(base, p)` |
| `path.normalize()` | `filepath.Clean(p)` |
| `Files.walkFileTree` | `filepath.WalkDir` |

### encoding/json

| Java | Go |
|---|---|
| Jackson `ObjectMapper.writeValueAsString` | `json.Marshal` |
| `readValue(json, User.class)` | `json.Unmarshal(data, &user)` |
| `@JsonProperty("name")` | `` `json:"name"` `` |
| `@JsonIgnore` | `` `json:"-"` `` |
| `@JsonInclude(NON_EMPTY)` | `omitempty` |
| `JsonNode` | `any` / `map[string]any` / `json.RawMessage` |
| `FAIL_ON_UNKNOWN_PROPERTIES` | `Decoder.DisallowUnknownFields()` |

### encoding/csv

| Java | Go |
|---|---|
| OpenCSV `CSVReader` | `csv.Reader` |
| Apache Commons CSV `CSVParser` | `csv.Reader` + a loop |
| `CSVWriter.close()` | `w.Flush()` then `w.Error()` |

### encoding/xml

| Java | Go |
|---|---|
| JAXB `@XmlRootElement` | `XMLName xml.Name` field |
| `@XmlAttribute` | `xml:"id,attr"` |
| `@XmlValue` | `xml:",chardata"` |
| StAX `XMLStreamReader` | `xml.Decoder.Token` |

### encoding/gob

| Java | Go |
|---|---|
| `ObjectOutputStream`/`ObjectInputStream` (`Serializable`) | `gob.Encoder` / `gob.Decoder` |
| `serialVersionUID` | None: matching is by field name |

### regexp

| Java | Go |
|---|---|
| `Pattern.compile(re)` | `regexp.MustCompile(re)` |
| `matcher.matches()` | `re.MatchString(s)` with `^...$` |
| `matcher.find()` in a loop | `re.FindAllString(s, -1)` |
| `matcher.group(1)` | `re.FindStringSubmatch(s)[1]` |
| `replaceAll("$1")` | `re.ReplaceAllString(s, "$1")` |
| `(?<name>...)` | `(?P<name>...)` |

### text/template

| Java | Go |
|---|---|
| Thymeleaf / FreeMarker / Mustache | `text/template` (text) or `html/template` (HTML) |
| `${user.name}` | `{{.User.Name}}` |
| `th:each` | `{{range}}` |
| `th:if` | `{{if}}` |
| `th:fragment` + `th:replace` | `{{define}}` + `{{template}}` |

### html/template

| Java | Go |
|---|---|
| Thymeleaf `th:text` (escaped) | `{{.}}` (escaped by default) |
| Thymeleaf `th:utext` (unescaped) | `template.HTML(...)` |
| JSP `<c:out>` | automatic contextual escaping |

### log/slog

| Java | Go |
|---|---|
| SLF4J `log.info("x {}", v)` | `slog.Info("x", "v", v)` |
| MDC | `logger.With(...)` / `slog.Group` |
| Logback JSON encoder | `slog.NewJSONHandler` |
| Log levels (`TRACE..ERROR`) | `slog.LevelDebug/Info/Warn/Error` |

### sync

| Java | Go |
|---|---|
| `synchronized`, `ReentrantLock` | `sync.Mutex` (not reentrant) |
| `ReadWriteLock` | `sync.RWMutex` |
| `CountDownLatch` | `sync.WaitGroup` |
| `ConcurrentHashMap` | `map` + `sync.Mutex`, or `sync.Map` |
| `Condition` | `sync.Cond` |
| Holder/lazy singleton | `sync.Once` / `sync.OnceValue` |

### sync/atomic

| Java | Go |
|---|---|
| `AtomicInteger` / `AtomicLong` | `atomic.Int32` / `atomic.Int64` |
| `AtomicBoolean` | `atomic.Bool` |
| `AtomicReference<T>` | `atomic.Pointer[T]` |
| `volatile` | No equivalent: use `atomic` or a mutex |

### context

| Java | Go |
|---|---|
| `Future.cancel`, `Thread.interrupt` | `cancel()` + `<-ctx.Done()` |
| `CompletableFuture.orTimeout` | `context.WithTimeout` |
| `ThreadLocal`, MDC | `context.WithValue` |
| Spring request attributes | values on `r.Context()` |

### channels, select and concurrency patterns

| Java | Go |
|---|---|
| `BlockingQueue` | buffered channel |
| `ExecutorService` fixed pool | N goroutines reading one jobs channel |
| `Semaphore` | buffered `chan struct{}` |
| `CompletableFuture.anyOf` | `select` or a result channel |
| Reactive streams | pipeline of channels |

### net/http

| Java | Go |
|---|---|
| `@RestController` method | `http.Handler` / `HandlerFunc` |
| `@GetMapping("/items/{id}")` | `mux.HandleFunc("GET /items/{id}", h)` |
| `@PathVariable` | `r.PathValue("id")` |
| `@RequestParam` | `r.URL.Query().Get("q")` |
| `@RequestBody` | `json.NewDecoder(r.Body).Decode(&v)` |
| `ResponseEntity.status(201)` | `w.WriteHeader(201)` |
| `Filter` / `HandlerInterceptor` | middleware `func(http.Handler) http.Handler` |
| `RestTemplate` / `WebClient` | `http.Client` |
| embedded Tomcat | `http.Server` |

### net/http/httptest

| Java | Go |
|---|---|
| Spring `MockMvc` | `httptest.NewRecorder` + `ServeHTTP` |
| `@SpringBootTest(webEnvironment = RANDOM_PORT)` | `httptest.NewServer` |
| WireMock / MockWebServer | `httptest.NewServer` or a fake `RoundTripper` |
| `TestRestTemplate` | `srv.Client()` |

### net

| Java | Go |
|---|---|
| `ServerSocket` / `Socket` | `net.Listener` / `net.Conn` |
| `InetAddress.getByName` | `net.ParseIP` / `net.LookupHost` |
| `socket.setSoTimeout` | `conn.SetReadDeadline` |
| `DatagramSocket` | `net.ListenPacket("udp", ...)` |
| Netty / NIO | goroutine per connection on `net.Conn` |

### net/url

| Java | Go |
|---|---|
| `java.net.URI` / `URL` | `*url.URL` |
| `UriComponentsBuilder` | `url.URL` + `url.Values` |
| `URLEncoder.encode(s, UTF_8)` | `url.QueryEscape(s)` |
| `URI.resolve` | `base.ResolveReference(ref)` |

### crypto/tls

| Java | Go |
|---|---|
| `SSLContext`, `SSLSocket` | `tls.Config`, `tls.Conn` |
| `KeyStore` / `TrustStore` | `tls.Certificate` / `x509.CertPool` |
| `HttpsURLConnection` hostname verifier | `ServerName` + built-in verification |
| Spring Boot `server.ssl.*` | `http.Server{TLSConfig}` + `ServeTLS` |
| mutual TLS (`clientAuth=need`) | `ClientAuth: tls.RequireAndVerifyClientCert` |

### database/sql

| Java | Go |
|---|---|
| `DataSource` / HikariCP | `*sql.DB` (pool built in) |
| `Connection.prepareStatement` | `db.Prepare` / placeholders in `Query` |
| `ResultSet.next()` + `getString` | `rows.Next()` + `rows.Scan(&s)` |
| try-with-resources | `defer rows.Close()` |
| `@Transactional` | `tx, _ := db.BeginTx(...)` + `defer tx.Rollback()` + `tx.Commit()` |
| `JdbcTemplate` / JPA | not in stdlib: see `sqlx`, `pgx`, `sqlc`, `GORM` |

### testing

| Java | Go |
|---|---|
| JUnit `@Test` | `func TestXxx(t *testing.T)` |
| `@ParameterizedTest` | table-driven test + `t.Run` |
| AssertJ / Hamcrest | plain `if got != want { t.Errorf(...) }` (or `testify`) |
| `@BeforeEach` / `@AfterEach` | setup in the test + `t.Cleanup` |
| `@BeforeAll` | `TestMain` |
| `@TempDir` | `t.TempDir()` |
| Mockito | small interfaces + hand-written fakes |
| JMH | `testing.B` benchmarks |

### reflect

| Java | Go |
|---|---|
| `Class<?>`, `getClass()` | `reflect.TypeOf(x)` |
| `getDeclaredFields()` | `t.Field(i)` for `i < t.NumField()` |
| annotations (`@NotNull`) | struct tags (`` `validate:"required"` ``) |
| `Method.invoke(obj, args)` | `v.MethodByName("X").Call(args)` |
| `Field.set(obj, v)` + `setAccessible` | `v.Elem().Field(i).Set(...)` (exported fields only) |

### unsafe

| Java | Go |
|---|---|
| `sun.misc.Unsafe` / `VarHandle` | `unsafe.Pointer` |
| `ByteBuffer` views over memory | `unsafe.Slice(ptr, n)` |
| `Unsafe.objectFieldOffset` | `unsafe.Offsetof` |

### generics, cmp and slices

| Java | Go |
|---|---|
| `<T extends Comparable<T>>` | `[T cmp.Ordered]` |
| `List<T>` | `[]T` or `type List[T any] ...` |
| wildcards `? extends T` | no wildcards: use constraints or interfaces |
| type erasure | no erasure: separate instantiations (`Stack[int]` is a real type) |
| `Collections.max` | `slices.Max` |

### runtime and runtime/debug

| Java | Go |
|---|---|
| `Thread.activeCount()` | `runtime.NumGoroutine()` |
| `Runtime.availableProcessors()` | `runtime.NumCPU()` |
| `Runtime.totalMemory()/freeMemory()` | `runtime.ReadMemStats` |
| `System.gc()` | `runtime.GC()` |
| `Thread.dumpStack()` | `debug.PrintStack()` / `runtime.Stack` |
| `StackWalker` | `runtime.Caller` / `runtime.Callers` |
| JVM flags `-Xmx` | `GOMEMLIMIT`, `GOGC` |

### runtime/pprof and net/http/pprof

| Java | Go |
|---|---|
| JFR / VisualVM / async-profiler | `go tool pprof` + `runtime/pprof` |
| Thread dump (`jstack`) | `/debug/pprof/goroutine?debug=2` |
| Heap dump + MAT | `/debug/pprof/heap` (sampled, not a full dump) |
| Spring Boot Actuator | `net/http/pprof` on an admin port |

### runtime/trace

| Java | Go |
|---|---|
| JFR events | `trace.Log` / regions + `go tool trace` |
| OpenTelemetry spans | `trace.NewTask` is local to one process; use OpenTelemetry for distributed tracing |
| Thread timeline views | the goroutine and processor timelines in `go tool trace` |

### crypto/*

| Java | Go |
|---|---|
| `MessageDigest.getInstance("SHA-256")` | `sha256.New()` / `sha256.Sum256` |
| `Mac.getInstance("HmacSHA256")` | `hmac.New(sha256.New, key)` |
| `SecureRandom` | `crypto/rand` |
| `Cipher.getInstance("AES/GCM/NoPadding")` | `cipher.NewGCM(block)` |
| `Signature` (Ed25519) | `ed25519.Sign` / `Verify` |
| `MessageDigest.isEqual` | `subtle.ConstantTimeCompare` |
| Spring Security `BCryptPasswordEncoder` | `golang.org/x/crypto/bcrypt` |

### embed

| Java | Go |
|---|---|
| `classpath:` resources / `getResourceAsStream` | `//go:embed` + `embed.FS` |
| Spring Boot fat JAR with static assets | single binary with embedded assets |
| Flyway `db/migration` scripts on the classpath | embedded `*.sql` files |
| Thymeleaf `templates/` directory | embedded `templates/*.tmpl` |

### flag

| Java | Go |
|---|---|
| picocli / Apache Commons CLI | `flag` (simple) or `cobra` (large CLIs) |
| `@Option(names="-v")` | `fs.Bool("v", false, "verbose")` |
| `@Parameters` | `fs.Args()` |
| Spring `@Value("${...}")` | flags + `os.Getenv` |

### os/exec

| Java | Go |
|---|---|
| `ProcessBuilder` / `Runtime.exec` | `exec.Command` |
| `Process.waitFor()` | `cmd.Wait()` / `cmd.Run()` |
| `process.getInputStream()` | `cmd.StdoutPipe()` or `cmd.Output()` |
| `Process.destroy()` | `cmd.Process.Kill()` or `CommandContext` |
| `pb.environment()` | `cmd.Env` |

### os/signal

| Java | Go |
|---|---|
| `Runtime.addShutdownHook(Thread)` | `signal.NotifyContext` + cleanup after `<-ctx.Done()` |
| Spring Boot graceful shutdown (`server.shutdown=graceful`) | `srv.Shutdown(ctx)` |
| `@PreDestroy` | `defer` and explicit cleanup in `main` |
| `kill -15` / `SIGTERM` | the same signal; handle it explicitly |
